<script lang="ts">
	import * as Card from '$lib/components/ui/card/index.js';
	import type { HostsRecord } from '$lib/pocketbase-types';
	import { cn } from '$lib/utils';
	import { Trash2, BellRing } from 'lucide-svelte';
	import { Button } from './ui/button';
	import { hostsStore } from '$lib/stores/hosts';

	let { host, class: className }: { host: HostsRecord; class?: string } = $props();

	function wake() {
		hostsStore.wakeHost(host);
	}

	function deleteHost() {
		if (window.confirm('Are you sure you want to delete this host?')) {
			hostsStore.deleteHost(host);
		}
	}
</script>

<Card.Root class={cn('relative', className)}>
	<Card.Content>
		<Button
			size="icon"
			class="absolute right-2 top-2 h-6 w-6 rounded-full"
			variant="destructive"
			onclick={deleteHost}
		>
			<Trash2 class="h-3 w-3" />
		</Button>
		<div class={cn('grid grid-cols-2 gap-2', className)}>
			<p class="text-md font-bold">
				Name: <span class="font-mono font-medium">{host.name}</span>
			</p>
			<p class="text-md font-bold">
				Broadcast IP: <span class="font-mono font-medium">{host.ip}</span>
			</p>
			<p class="text-md font-bold">
				Mac Address: <span class="font-mono font-medium">{host.mac}</span>
			</p>
			<p class="text-md font-bold">
				Port: <span class="font-mono font-medium">{host.port}</span>
			</p>
		</div>
	</Card.Content>
	<Card.Footer class="flex justify-between gap-2">
		<Button
			size="sm"
			variant="outline"
			class="w-full bg-green-400/60 hover:bg-green-400/40"
			onclick={wake}
		>
			<BellRing class="h-4 w-4" />
		</Button>
	</Card.Footer>
</Card.Root>
