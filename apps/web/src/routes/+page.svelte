<script lang="ts">
	import { goto } from '$app/navigation';
	import CreateHostForm from '$lib/components/CreateHostForm.svelte';
	import HostCard from '$lib/components/HostCard.svelte';
	import { Button } from '$lib/components/ui/button/index.js';
	import { pb } from '$lib/pb';
	import { hostsStore } from '$lib/stores/hosts';
	import autoAnimate from '@formkit/auto-animate';
	import { onMount } from 'svelte';

	$effect(() => {
		if (pb.authStore.isValid) {
			hostsStore.fetchHosts();
		} else {
			goto('/auth');
		}
	});
</script>

<main class="pt-20">
	<div class="space-y-2">
		<CreateHostForm class="mx-auto max-w-[40em]" />
		<ul use:autoAnimate class="space-y-2">
			{#each $hostsStore as host (host.id)}
				<HostCard {host} class="mx-auto max-w-[40em]" />
			{/each}
		</ul>
	</div>
</main>
