<script lang="ts">
	import type { PageProps } from './$types';
	import SearchComponent from '$lib/components/SearchComponent.svelte';

	let { data }: PageProps = $props();
</script>

<svelte:head>
	<title>Search Notes</title>
</svelte:head>

<div class="space-y-8">
	<div class="relative">
		<div class="absolute inset-0 flex items-center" aria-hidden="true">
			<div class="w-full border-t border-gray-200 dark:border-gray-800"></div>
		</div>
		<div class="relative flex justify-center">
			<span class="px-4 text-sm text-gray-500 bg-white dark:bg-gray-950 dark:text-gray-400"
				>Search</span
			>
		</div>
	</div>

	<div class="max-w-2xl mx-auto text-center">
		<h1 class="text-4xl font-bold tracking-tight text-gray-900 dark:text-white">Find your notes</h1>
		<p class="mt-4 text-gray-600 dark:text-gray-400">
			Search through all my published notes by title, content, or tags.
		</p>
	</div>

	<div class="max-w-3xl mx-auto mt-8">
		<SearchComponent notes={data.notes} />
	</div>
</div>
