<script lang="ts">
	import { onMount } from 'svelte';
	import { Search, Loader2, FileText, Tag, Calendar } from 'lucide-svelte';
	import type { Note } from '$lib/api';
	import lunr from 'lunr';

	const { notes = [] } = $props();

	let searchQuery = $state('');
	let searchResults = $state<Array<Note & { score?: number }>>([]);
	let searchIndex = $state<lunr.Index | undefined>(undefined);
	let loading = $state(true);
	let searching = $state(false);
	let selectedResult = $state(-1);

	function handleKeyDown(event: KeyboardEvent) {
		if (searchResults.length === 0) return;

		if (event.key === 'ArrowDown') {
			event.preventDefault();
			selectedResult = (selectedResult + 1) % searchResults.length;
		} else if (event.key === 'ArrowUp') {
			event.preventDefault();
			selectedResult = selectedResult <= 0 ? searchResults.length - 1 : selectedResult - 1;
		} else if (event.key === 'Enter' && selectedResult >= 0) {
			event.preventDefault();
			window.location.href = `/note/${searchResults[selectedResult].id}`;
		}
	}

	onMount(async () => {
		try {
			searchIndex = lunr(function (this: any) {
				this.field('title', { boost: 10 });
				this.field('description', { boost: 5 });
				this.field('content');
				this.field('tags', { boost: 5 });
				this.ref('id');

				notes.forEach((note) => {
					const { id, metadata, content } = note;
					this.add({
						id,
						title: metadata?.title || id,
						description: metadata?.description || '',
						content: content,
						tags: metadata?.tags ? metadata.tags.join(' ') : ''
					});
				});
			});
			loading = false;
		} catch (error) {
			console.error('Error creating search index:', error);
			loading = false;
		}
	});

	async function handleSearch() {
		if (!searchQuery.trim()) {
			searchResults = [];
			return;
		}

		searching = true;
		selectedResult = -1;

		try {
			if (!searchIndex) {
				searchResults = [];
				return;
			}

			const results = searchIndex.search(searchQuery);
			searchResults = results.map((result: any) => {
				const note = notes.find((n) => n.id === result.ref);
				return {
					...note,
					score: result.score
				};
			});
		} catch (error) {
			console.error('Search error:', error);
			searchResults = [];
		} finally {
			searching = false;
		}
	}

	function highlightMatch(text: string, query: string) {
		if (!query || !text) return text;

		const regex = new RegExp(`(${query.replace(/[.*+?^${}()|[\]\\]/g, '\\$&')})`, 'gi');
		return text.replace(
			regex,
			'<mark class="bg-yellow-100 dark:bg-yellow-900/50 text-gray-900 dark:text-gray-100 px-0.5 rounded">$1</mark>'
		);
	}

	$effect(() => {
		if (searchQuery) {
			handleSearch();
		}
	});
</script>

<div class="space-y-6">
	<div class="relative">
		<div class="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
			{#if searching}
				<Loader2 class="h-5 w-5 animate-spin text-gray-400" />
			{:else}
				<Search class="h-5 w-5 text-gray-400" />
			{/if}
		</div>
		<input
			type="search"
			class="block w-full rounded-lg border border-gray-300 bg-white p-4 pl-10 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-700 dark:bg-gray-800 dark:text-white dark:placeholder-gray-400"
			placeholder="Search by title, content, or tags..."
			bind:value={searchQuery}
			onkeydown={handleKeyDown}
		/>
	</div>

	{#if loading}
		<div class="flex items-center justify-center py-8">
			<Loader2 class="h-8 w-8 animate-spin text-blue-500" />
			<span class="ml-2 text-gray-600 dark:text-gray-400">Loading search index...</span>
		</div>
	{:else if searchQuery && searchResults.length === 0}
		<div class="py-8 text-center">
			<p class="text-gray-600 dark:text-gray-400">No results found for "{searchQuery}"</p>
		</div>
	{:else if searchResults.length > 0}
		<div class="space-y-4">
			<p class="text-sm text-gray-500 dark:text-gray-400">{searchResults.length} results found</p>
			{#each searchResults as result, i}
				<a
					href={`/note/${result.id}`}
					class="block rounded-lg border border-gray-200 bg-white p-4 transition-colors hover:bg-gray-50 dark:border-gray-800 dark:bg-gray-900 dark:hover:bg-gray-800"
					class:ring-2={i === selectedResult}
					class:ring-blue-500={i === selectedResult}
				>
					<div class="flex items-start">
						<FileText class="mr-3 mt-0.5 h-5 w-5 flex-shrink-0 text-blue-500 dark:text-blue-400" />
						<div class="min-w-0 flex-1">
							<h3 class="truncate text-lg font-semibold text-gray-900 dark:text-white">
								{@html highlightMatch(result.metadata?.title || result.id, searchQuery)}
							</h3>
							{#if result.metadata?.description}
								<p class="mt-1 line-clamp-2 text-gray-600 dark:text-gray-400">
									{@html highlightMatch(result.metadata.description, searchQuery)}
								</p>
							{/if}
							<div class="mt-2 flex flex-wrap gap-2">
								{#if result.metadata?.tags && result.metadata.tags.length > 0}
									<div class="mr-4 flex items-center text-xs text-gray-500 dark:text-gray-400">
										<Tag class="mr-1 h-3.5 w-3.5" />
										<span>
											{result.metadata.tags.slice(0, 3).join(', ')}
											{#if result.metadata.tags.length > 3}
												<span class="text-gray-400 dark:text-gray-500"
													>+{result.metadata.tags.length - 3}</span
												>
											{/if}
										</span>
									</div>
								{/if}
								{#if result.metadata?.updated}
									<div class="flex items-center text-xs text-gray-500 dark:text-gray-400">
										<Calendar class="mr-1 h-3.5 w-3.5" />
										<span>{new Date(result.metadata.updated).toLocaleDateString()}</span>
									</div>
								{/if}
							</div>
						</div>
					</div>
				</a>
			{/each}
		</div>
	{/if}
</div>
