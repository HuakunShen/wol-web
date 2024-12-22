import { PUBLIC_PB_URL } from '$env/static/public';
import { pb } from '$lib/pb';
import type { HostsRecord, RecordIdString } from '$lib/pocketbase-types';
import { toast } from 'svelte-sonner';
import { writable } from 'svelte/store';

export function createHostsStore() {
	const hosts = writable<HostsRecord[]>([]);

	async function fetchHosts() {
		const records = await pb.collection('hosts').getFullList();
		hosts.set(records);
	}

	async function createHost(host: {
		ip: string;
		mac: string;
		name: string;
		port: number;
		user: RecordIdString;
	}) {
		await pb.collection('hosts').create(host);
		fetchHosts();
	}

	async function updateHost(host: HostsRecord) {
		await pb.collection('hosts').update(host.id, host);
		fetchHosts();
	}

	async function deleteHost(host: HostsRecord) {
		await pb.collection('hosts').delete(host.id);
		fetchHosts();
	}

	async function wakeHost(host: HostsRecord) {
		pb.send('/api/wake', {
			method: 'POST',
			headers: {
				Accept: 'application/json',
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ id: host.id })
		})
			.then((res) => {
				toast.success('WakeOnLan Magic Packet Sent');
			})
			.catch((err) => {
				toast.error('Failed to wake host', {
					description: err.message
				});
			});
	}

	return {
		...hosts,
		fetchHosts,
		createHost,
		updateHost,
		deleteHost,
		wakeHost
	};
}

export const hostsStore = createHostsStore();
