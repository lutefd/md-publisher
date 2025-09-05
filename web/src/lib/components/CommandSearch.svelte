<script lang="ts">
	import type { Note } from '$lib/api';
	import { onMount } from 'svelte';
	import { Search } from 'lucide-svelte';
	import { Command, Dialog } from 'bits-ui';
	import lunr from 'lunr';

	const { notes = [] } = $props();
	let localNotes: Note[] = $state(notes);

	let open = $state(false);
	let searchQuery = $state('');
	let searchResults = $state<Array<Note & { score?: number }>>([]);
	let searchIndex = $state<lunr.Index | undefined>(undefined);
	let loading = $state(true);

	function recomputeSearch() {
		try {
			if (!searchIndex) return;
			if (!searchQuery.trim()) {
				searchResults = [];
				return;
			}

			const results = searchIndex.search(searchQuery);
			const processedResults = [];
			for (const result of results) {
				if (!result || !result.ref) continue;
				const note = localNotes.find((n) => n && n.id === result.ref);
				if (note && note.metadata) {
					processedResults.push({ ...note, score: result.score });
				}
			}
			searchResults = processedResults;
		} catch (err) {
			console.error('Search recompute error:', err);
			searchResults = [];
		}
	}

	$effect(() => {
		if (!open) {
			searchQuery = '';
			searchResults = [];
		}
	});

	$effect(() => {
		void searchIndex;
		void searchQuery;
		recomputeSearch();
	});

	function handleGlobalKeyDown(event: KeyboardEvent) {
		if ((event.metaKey || event.ctrlKey) && event.key === 'k') {
			event.preventDefault();
			open = true;
		}
	}

	function handleCommandKeyDown(event: KeyboardEvent) {
		if (event.key === 'Escape') {
			open = false;
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

				localNotes.forEach((note) => {
					if (!note || !note.id) return;

					const { id, metadata, content } = note;
					if (!metadata) return;

					this.add({
						id,
						title: metadata.title || id,
						description: metadata.description || '',
						content: content || '',
						tags: metadata.tags ? metadata.tags.join(' ') : ''
					});
				});
			});

			loading = false;
			recomputeSearch();
		} catch (error) {
			console.error('Error creating search index:', error);
			loading = false;
		}
	});

	function handleSelect(result: Note & { score?: number }) {
		if (result && result.id) {
			window.location.href = `/note/${result.id}`;
			open = false;
		}
	}
</script>

<svelte:window onkeydown={handleGlobalKeyDown} />

<button
	onclick={() => (open = true)}
	class="inline-flex items-center justify-center px-3 text-sm font-medium transition-colors bg-white border border-gray-200 rounded-md h-9 ring-offset-white hover:bg-gray-100 hover:text-gray-900 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-gray-950 focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 dark:border-gray-800 dark:bg-gray-950 dark:ring-offset-gray-950 dark:hover:bg-gray-800 dark:hover:text-gray-50 dark:focus-visible:ring-gray-300"
>
	<Search class="w-4 h-4 mr-2" />
	<span class="sr-only sm:not-sr-only sm:whitespace-nowrap">Search</span>
	<kbd
		class="pointer-events-none ml-2 hidden h-5 select-none items-center gap-1 rounded border border-gray-200 bg-gray-100 px-1.5 font-mono text-[10px] font-medium text-gray-600 opacity-100 sm:flex dark:border-gray-800 dark:bg-gray-900 dark:text-gray-400"
	>
		<span class="text-xs">⌘</span>K
	</kbd>
</button>

<Dialog.Root
	bind:open
	onOpenChange={(isOpen: boolean) => {
		if (isOpen) {
			setTimeout(() => {
				const input = document.querySelector('.command-input') as HTMLElement;
				if (input && 'focus' in input) input.focus();
			}, 0);
		}
	}}
>
	<Dialog.Portal>
		<Dialog.Overlay
			class="data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 fixed inset-0 z-50 bg-black/50 backdrop-blur-sm"
		/>
		<Dialog.Content
			class="data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[state=closed]:slide-out-to-left-1/2 data-[state=closed]:slide-out-to-top-[48%] data-[state=open]:slide-in-from-left-1/2 data-[state=open]:slide-in-from-top-[48%] fixed left-[50%] top-[50%] z-50 grid w-full max-w-lg translate-x-[-50%] translate-y-[-50%] gap-4 border border-gray-200 bg-white p-0 shadow-lg duration-200 sm:rounded-lg dark:border-gray-800 dark:bg-gray-950"
		>
			<Command.Root
				class="flex flex-col overflow-hidden rounded-md"
				onkeydown={handleCommandKeyDown}
				shouldFilter={false}
			>
				<Command.Input
					class="flex w-full h-12 px-4 py-3 text-sm bg-transparent border-0 border-b border-gray-200 rounded-md outline-none command-input placeholder:text-gray-500 focus:ring-0 disabled:cursor-not-allowed disabled:opacity-50 dark:border-gray-800 dark:placeholder:text-gray-400"
					bind:value={searchQuery}
					placeholder="Search notes..."
				/>
				<Command.List class="max-h-[300px] overflow-y-auto overflow-x-hidden p-2">
					{#if loading}
						<Command.Loading>
							<div class="p-4 text-sm text-center text-gray-500 dark:text-gray-400">Loading...</div>
						</Command.Loading>
					{:else if !searchQuery.trim()}
						<div class="p-4 text-sm text-center text-gray-500 dark:text-gray-400">
							Start typing to search notes...
						</div>
					{:else if searchQuery && searchResults.length === 0}
						<Command.Empty>
							<div class="p-4 text-sm text-center text-gray-500 dark:text-gray-400">
								No results found for "{searchQuery}"
							</div>
						</Command.Empty>
					{:else if searchResults.length > 0}
						<Command.Group>
							{#each searchResults as result, i}
								{#if result && result.metadata}
									<Command.Item
										value={`${result.id} ${result.metadata?.title ?? ''} ${result.metadata?.description ?? ''} ${(result.metadata?.tags ?? []).join(' ')}`}
										onSelect={() => handleSelect(result)}
										class="relative flex cursor-pointer select-none items-center rounded-md px-2 py-2.5 text-sm outline-none hover:bg-gray-100 aria-selected:bg-gray-100 dark:hover:bg-gray-800 dark:aria-selected:bg-gray-800"
									>
										<div class="flex-1 truncate">
											<h4 class="font-medium text-gray-900 truncate dark:text-white">
												{result.metadata.title || result.id}
											</h4>
											{#if result.metadata.description}
												<p class="text-xs text-gray-500 truncate dark:text-gray-400">
													{result.metadata.description}
												</p>
											{/if}
										</div>
									</Command.Item>
								{/if}
							{/each}
						</Command.Group>
					{/if}
				</Command.List>
			</Command.Root>
			<div
				class="flex items-center justify-between p-2 border-t border-gray-200 dark:border-gray-800"
			>
				<div class="text-xs text-gray-500 dark:text-gray-400">
					Press <kbd
						class="px-1 text-xs bg-gray-100 border border-gray-200 rounded dark:border-gray-800 dark:bg-gray-900"
						>↑</kbd
					>
					<kbd
						class="px-1 text-xs bg-gray-100 border border-gray-200 rounded dark:border-gray-800 dark:bg-gray-900"
						>↓</kbd
					> to navigate
				</div>
				<div class="text-xs text-gray-500 dark:text-gray-400">
					Press <kbd
						class="px-1 text-xs bg-gray-100 border border-gray-200 rounded dark:border-gray-800 dark:bg-gray-900"
						>Enter</kbd
					> to select
				</div>
			</div>
		</Dialog.Content>
	</Dialog.Portal>
</Dialog.Root>
