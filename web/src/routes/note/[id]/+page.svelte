<script lang="ts">
	import type { PageProps } from './$types';
	import { onMount } from 'svelte';
	import MetadataDisplay from '$lib/components/MetadataDisplay.svelte';

	let { data }: PageProps = $props();
	let loading = false;

	onMount(() => {
		const backToTopButton = document.querySelector('.back-to-top');
		if (backToTopButton) {
			backToTopButton.addEventListener('click', () => {
				window.scrollTo({ top: 0, behavior: 'smooth' });
			});
		}
	});
</script>

<svelte:head>
	<title>{data.note?.metadata?.title || data.note?.id}</title>
</svelte:head>

<article class="mx-auto">
	<div class="mb-10 space-y-4">
		<div class="flex items-center space-x-2 text-sm">
			<a href="/" class="flex items-center text-blue-600 hover:underline dark:text-blue-400">
				<svg
					xmlns="http://www.w3.org/2000/svg"
					class="mr-1 h-4 w-4"
					viewBox="0 0 24 24"
					fill="none"
					stroke="currentColor"
					stroke-width="2"
					stroke-linecap="round"
					stroke-linejoin="round"><path d="m15 18-6-6 6-6" /></svg
				>
				Back to notes
			</a>
			<span class="text-gray-400 dark:text-gray-600">/</span>
			<span class="text-gray-600 dark:text-gray-400">{data.note?.id}</span>
		</div>

		{#if loading}
			<div class="py-8 text-center">
				<div
					class="inline-block h-8 w-8 animate-spin rounded-full border-b-2 border-t-2 border-blue-500"
				></div>
				<p class="mt-2 text-gray-600 dark:text-gray-400">Loading note...</p>
			</div>
		{:else if data.note}
			<h1 class="text-4xl font-bold text-gray-900 dark:text-white">
				{data.note.metadata?.title || data.note.id}
			</h1>

			{#if data.note.metadata?.description}
				<p class="text-xl text-gray-600 dark:text-gray-400">{data.note.metadata.description}</p>
			{/if}

			{#if data.note.metadata?.tags && data.note.metadata.tags.length > 0}
				<div class="flex flex-wrap gap-2">
					{#each data.note.metadata.tags as tag}
						<span
							class="inline-flex items-center rounded-full bg-blue-50 px-2.5 py-1 text-xs font-medium text-blue-700 dark:bg-blue-900/30 dark:text-blue-300"
						>
							{tag}
						</span>
					{/each}
				</div>
			{/if}

			<!-- Display additional metadata from frontmatter -->
			<MetadataDisplay metadata={data.note.metadata || {}} />

			<div class="border-b border-gray-200 pt-4 dark:border-gray-800"></div>
		{/if}
	</div>

	<div class="prose prose-gray dark:prose-invert max-w-none">
		{#if loading}
			<!-- Loading placeholder already shown above -->
		{:else if data.note}
			{@html data.content}
		{:else}
			<p>Note not found.</p>
		{/if}
	</div>

	<div class="mt-16 border-t border-gray-200 pt-8 dark:border-gray-800">
		<div class="flex items-center justify-between">
			<a
				href="/"
				class="inline-flex h-9 items-center justify-center rounded-md border border-gray-200 bg-white px-4 py-2 text-sm font-medium ring-offset-white transition-colors hover:bg-gray-100 hover:text-gray-900 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-gray-950 focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 dark:border-gray-800 dark:bg-gray-950 dark:ring-offset-gray-950 dark:hover:bg-gray-800 dark:hover:text-gray-50 dark:focus-visible:ring-gray-300"
			>
				<svg
					xmlns="http://www.w3.org/2000/svg"
					class="mr-2 h-4 w-4"
					viewBox="0 0 24 24"
					fill="none"
					stroke="currentColor"
					stroke-width="2"
					stroke-linecap="round"
					stroke-linejoin="round"><path d="m15 18-6-6 6-6" /></svg
				>
				Back to notes
			</a>
			<button
				class="back-to-top inline-flex h-9 items-center justify-center rounded-md border border-gray-200 bg-white px-4 py-2 text-sm font-medium ring-offset-white transition-colors hover:bg-gray-100 hover:text-gray-900 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-gray-950 focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 dark:border-gray-800 dark:bg-gray-950 dark:ring-offset-gray-950 dark:hover:bg-gray-800 dark:hover:text-gray-50 dark:focus-visible:ring-gray-300"
			>
				<svg
					xmlns="http://www.w3.org/2000/svg"
					class="mr-2 h-4 w-4"
					viewBox="0 0 24 24"
					fill="none"
					stroke="currentColor"
					stroke-width="2"
					stroke-linecap="round"
					stroke-linejoin="round"><path d="m18 15-6-6-6 6" /></svg
				>
				Back to top
			</button>
		</div>
	</div>
</article>
