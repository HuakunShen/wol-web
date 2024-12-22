<script lang="ts">
	import { superForm, defaults } from 'sveltekit-superforms';
	import SuperDebug from 'sveltekit-superforms';
	import { zodClient, zod } from 'sveltekit-superforms/adapters';
	import { z } from 'zod';
	import * as Form from '$lib/components/ui/form/index.js';
	import { Input } from './ui/input';
	import { toast } from 'svelte-sonner';
	import { cn } from '$lib/utils';
	import { pb } from '$lib/pb';
	import { hostsStore } from '$lib/stores/hosts';

	const formSchema = z.object({
		name: z.string().min(1, 'Name is required'),
		mac: z
			.string()
			.min(1, 'MAC is required')
			.regex(/^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$/, 'Invalid MAC address format'),
		ip: z.string().min(1, 'IP is required').ip(),
		port: z.number().default(9)
	});

	const form = superForm(defaults(zod(formSchema)), {
		validators: zodClient(formSchema),
		SPA: true,
		onUpdate({ form, cancel }) {
			if (!form.valid) {
				toast.error('Invalid');
				return;
			}
			if (!pb.authStore.record) {
				toast.error('You must be logged in to create a host');
				return;
			}
			hostsStore.createHost({ ...form.data, user: pb.authStore.record?.id });
			toast.success('Host created');
			cancel();
		}
	});
	const { form: formData, enhance, errors } = form;

	let { class: className }: { class?: string } = $props();
</script>

<form method="POST" use:enhance class={cn('grid grid-cols-2 gap-2', className)}>
	<Form.Field {form} name="name">
		<Form.Control>
			{#snippet children({ props })}
				<div class="flex flex-col gap-2">
					<Form.Label>Name</Form.Label>
					<Input {...props} name="name" bind:value={$formData.name} placeholder="TrueNAS" />
				</div>
			{/snippet}
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>
	<Form.Field {form} name="mac">
		<Form.Control>
			{#snippet children({ props })}
				<div class="flex flex-col gap-2">
					<Form.Label>MAC</Form.Label>
					<Input {...props} name="mac" bind:value={$formData.mac} placeholder="86:2f:57:c1:df:65" />
				</div>
			{/snippet}
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>
	<Form.Field {form} name="ip">
		<Form.Control>
			{#snippet children({ props })}
				<div class="flex flex-col gap-2">
					<Form.Label>IP</Form.Label>
					<Input {...props} name="ip" bind:value={$formData.ip} placeholder="e.g. 192.168.1.1" />
				</div>
			{/snippet}
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>
	<Form.Field {form} name="port">
		<Form.Control>
			{#snippet children({ props })}
				<div class="flex flex-col gap-2">
					<Form.Label>Port</Form.Label>
					<Input {...props} name="port" bind:value={$formData.port} placeholder="Port" />
				</div>
			{/snippet}
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>
	<Form.Button class="my-1 w-full">Create</Form.Button>
	<SuperDebug data={$formData} />
</form>
