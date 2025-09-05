<script lang="ts">
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
</script>

<svelte:head>
	<title>Notes | Home</title>
</svelte:head>

<div class="space-y-8">
	<div class="relative">
		<div class="absolute inset-0 flex items-center" aria-hidden="true">
			<div class="w-full border-t border-gray-200 dark:border-gray-800"></div>
		</div>
		<div class="relative flex justify-center">
			<span class="bg-white px-4 text-sm text-gray-500 dark:bg-gray-950 dark:text-gray-400"
				>My Shared Knowledge Base</span
			>
		</div>
	</div>

	<h1 class="text-center text-4xl font-bold tracking-tight text-gray-900 dark:text-white">
		Welcome to my notes
	</h1>

	<p class="mx-auto max-w-2xl text-center text-gray-600 dark:text-gray-400">
		This is a self-hosted alternative to Obsidian Publish. Browse my published notes below or use
		the search feature to find specific content.
	</p>

	<div class="mt-10">
		{#if data.notes.length === 0}
			<div
				class="rounded-lg border border-dashed border-gray-300 p-12 text-center dark:border-gray-700"
			>
				<svg
					class="mx-auto h-12 w-12 text-gray-400 dark:text-gray-600"
					fill="none"
					viewBox="0 0 24 24"
					stroke="currentColor"
					stroke-width="1"
				>
					<path
						stroke-linecap="round"
						stroke-linejoin="round"
						d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10"
					></path>
				</svg>
				<h3 class="mt-2 text-sm font-semibold text-gray-900 dark:text-white">
					No notes published yet
				</h3>
				<p class="mt-1 text-sm text-gray-500 dark:text-gray-400">I'll publish my notes here</p>
			</div>
		{:else}
			<div class="grid gap-6 sm:grid-cols-2">
				{#each data.notes
					.slice(0, 6)
					.sort((a, b) => new Date(b.metadata.updated || '').getTime() - new Date(a.metadata.updated || '').getTime()) as note}
					<a
						href={`/note/${note.id}`}
						class="group relative flex flex-col overflow-hidden rounded-lg border border-gray-200 bg-white p-6 transition-all duration-200 hover:-translate-y-1 hover:shadow-md dark:border-gray-800 dark:bg-gray-900"
					>
						{#if note.metadata.tags && note.metadata.tags.length > 0}
							<div class="mb-3 flex flex-wrap gap-2">
								{#each note.metadata.tags.slice(0, 3) as tag}
									<span
										class="inline-flex items-center rounded-full bg-blue-50 px-2 py-1 text-xs font-medium text-blue-700 dark:bg-blue-900/30 dark:text-blue-300"
									>
										{tag}
									</span>
								{/each}
								{#if note.metadata.tags.length > 3}
									<span
										class="inline-flex items-center rounded-full bg-gray-50 px-2 py-1 text-xs font-medium text-gray-600 dark:bg-gray-800 dark:text-gray-400"
									>
										+{note.metadata.tags.length - 3}
									</span>
								{/if}
							</div>
						{/if}
						<h2
							class="text-xl font-semibold text-gray-900 transition-colors group-hover:text-blue-600 dark:text-white dark:group-hover:text-blue-400"
						>
							{note.metadata.title || note.id}
						</h2>
						{#if note.metadata.description}
							<p class="mt-2 line-clamp-3 text-gray-600 dark:text-gray-400">
								{note.metadata.description}
							</p>
						{/if}
						<div class="mt-4 flex items-center text-sm text-gray-500 dark:text-gray-400">
							<svg
								xmlns="http://www.w3.org/2000/svg"
								class="mr-1.5 h-4 w-4"
								viewBox="0 0 24 24"
								fill="none"
								stroke="currentColor"
								stroke-width="2"
								stroke-linecap="round"
								stroke-linejoin="round"
								><circle cx="12" cy="12" r="10" /><polyline points="12 6 12 12 16 14" /></svg
							>
							Updated {new Date(note.metadata.updated || '').toLocaleDateString()}
						</div>
					</a>
				{/each}
			</div>
		{/if}
	</div>
</div>
