<script lang="ts">
	import { Button, buttonVariants } from '$lib/components/ui/button/index.js';
	import * as Card from '$lib/components/ui/card/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Label } from '$lib/components/ui/label/index.js';
	import { Info } from 'lucide-svelte';
	import { pb } from '$lib/pb';
	import { toast } from 'svelte-sonner';
	import * as Popover from '$lib/components/ui/popover/index.js';
	import { goto } from '$app/navigation';
	import { dev } from '$app/environment';
	import { page } from '$app/state';

	let email = $state('');
	let password = $state('');

	$effect(() => {
		if (pb.authStore.isValid) {
			toast.warning('Already logged in');
			goto('/');
		}
	});

	function login(e: Event) {
		e.preventDefault();
		pb.collection('users')
			.authWithPassword(email, password)
			.then((authData) => {
				toast.success('Logged in', { description: authData.record.name });
				goto('/');
			})
			.catch((err) => {
				toast.error(err.message);
			});
	}
</script>

<div class="flex h-screen items-center justify-center">
	<Card.Root class="w-96">
		<form onsubmit={login}>
			<Card.Header>
				<Card.Title>Login</Card.Title>
				<Card.Description class="flex items-center justify-between"
					>Login to your account

					<Popover.Root>
						<Popover.Trigger class={buttonVariants({ variant: 'ghost', size: 'icon' })}>
							<Info />
						</Popover.Trigger>
						<Popover.Content class="w-80">
							{@const pbAdminUrl = dev
								? 'http://localhost:8090'
								: page.url.protocol + '//' + page.url.host + '/_/'}
							<Card.Description
								>Account can only be created by admin from <a target="_blank" href={pbAdminUrl}>
									{pbAdminUrl}
								</a>
							</Card.Description>
						</Popover.Content>
					</Popover.Root>
				</Card.Description>
			</Card.Header>
			<Card.Content>
				<div class="grid w-full items-center gap-4">
					<div class="flex flex-col space-y-1.5">
						<Label for="email">Email</Label>
						<Input id="email" placeholder="Email" bind:value={email} autofocus />
					</div>
					<div class="flex flex-col space-y-1.5">
						<Label for="password">Password</Label>
						<Input id="password" type="password" placeholder="Password" bind:value={password} />
					</div>
				</div>
			</Card.Content>
			<Card.Footer class="">
				<Button type="submit" class="w-full">Login</Button>
			</Card.Footer>
		</form>
	</Card.Root>
</div>
