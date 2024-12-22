<script lang="ts">
	import Navbar from '$lib/components/navbar.svelte';
	import '../app.css';
	import { ModeWatcher } from 'mode-watcher';
	import { Toaster } from 'svelte-sonner';
	let { children } = $props();
</script>

<ModeWatcher />
<Toaster richColors />
<Navbar />
{@render children?.()}
<footer class="bg-background fixed bottom-0 w-full px-3 py-2">
	<span class="text-foreground/60 text-xs"
		>Made with ❤️ by <a target="_blank" href="https://github.com/HuakunShen">Huakun</a></span
	>
</footer>
