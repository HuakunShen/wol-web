<script lang="ts">
	import * as Card from '$lib/components/ui/card/index.js';
	import type { HostsRecord } from '$lib/pocketbase-types';
	import { cn } from '$lib/utils';
	import { Trash2, BellRing } from 'lucide-svelte';
	import { Button } from './ui/button';
	import { hostsStore } from '$lib/stores/hosts';

	let { host, class: className }: { host: HostsRecord; class?: string } = $props();

	function wake(e: Event) {
		hostsStore.wakeHost(host);
	}

	function deleteHost(e: Event) {
		hostsStore.deleteHost(host);
	}
</script>

<Card.Root class={cn(className)}>
	<Card.Content>
		<div class={cn('grid grid-cols-2 gap-2', className)}>
			<p class="text-md font-bold">
				Name: <span class="font-mono font-medium">{host.name}</span>
			</p>
			<p class="text-md font-bold">
				IP Address: <span class="font-mono font-medium">{host.ip}</span>
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
		<Button size="sm" variant="outline" class="w-full bg-green-400/60" onclick={wake}>
			<BellRing class="h-4 w-4" />
		</Button>
		<Button size="sm" variant="destructive" class="w-full" onclick={deleteHost}>
			<Trash2 class="h-4 w-4" />
		</Button>
	</Card.Footer>
</Card.Root>
