<script>
	import { Button } from './ui/button';
	import { pb } from '$lib/pb';
	import { cn } from '$lib/utils';
	import { page } from '$app/state';
	import { LogOut } from 'lucide-svelte';
	import { goto } from '$app/navigation';

	function logout() {
		pb.authStore.clear();
		goto('/auth');
	}
</script>

<div
	class="bg-background fixed top-0 flex h-12 w-full items-center justify-between px-4 font-mono text-lg font-bold"
>
	<span></span>
	<div class="flex space-x-4">
		<a
			href="/"
			class={cn('hover:text-foreground/80 transition-colors', {
				'text-foreground/60': page.url.pathname != '/',
				'text-foreground': page.url.pathname == '/',
				underline: page.url.pathname == '/'
			})}
		>
			Home
		</a>
		<a
			href="/auth"
			class={cn('hover:text-foreground/80 transition-colors', {
				'text-foreground/60': page.url.pathname != '/auth',
				'text-foreground': page.url.pathname == '/auth',
				underline: page.url.pathname == '/auth'
			})}
		>
			Auth
		</a>
	</div>
	<Button size="icon" variant="outline" class="" onclick={logout}>
		<LogOut />
	</Button>
</div>
