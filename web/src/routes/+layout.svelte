<script lang="ts">
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';
	import CommandSearch from '$lib/components/CommandSearch.svelte';
	let { data, children } = $props();
	let notes = $derived(data.notes || []);
</script>

<svelte:head>
	<meta charset="UTF-8" />
	<meta name="description" content="Obsidian Publisher Alternative" />
	<meta name="viewport" content="width=device-width" />
	<link rel="icon" href={favicon} />
	<link rel="stylesheet" href="https://rsms.me/inter/inter.css" />
	<title>Notes</title>
</svelte:head>

<div class="flex flex-col min-h-screen">
	<header
		class="sticky top-0 z-40 w-full border-b border-gray-200 bg-white/80 backdrop-blur-sm dark:border-gray-800 dark:bg-gray-950/80"
	>
		<div class="container flex items-center h-16 px-4 sm:px-6 lg:px-8">
			<a href="/" class="flex items-center space-x-2">
				<span class="w-6 h-6 rounded-full bg-gradient-to-br from-blue-500 to-indigo-600"></span>
				<span class="text-xl font-bold text-gray-900 dark:text-white">Notes</span>
			</a>
			<div class="flex-1"></div>
			<nav class="flex items-center space-x-6">
				<a
					href="/"
					class="text-sm font-medium text-gray-600 transition-colors hover:text-gray-900 dark:text-gray-300 dark:hover:text-white"
					>Home</a
				>
				<a
					href="/search"
					class="text-sm font-medium text-gray-600 transition-colors hover:text-gray-900 dark:text-gray-300 dark:hover:text-white"
					>Search</a
				>
				<CommandSearch {notes} />
			</nav>
		</div>
	</header>
	<main class="container flex-1 max-w-4xl px-4 py-8 mx-auto sm:px-6 lg:px-8">
		{@render children?.()}
	</main>
	<footer class="py-6 border-t border-gray-200 dark:border-gray-800">
		<div class="container px-4 mx-auto sm:px-6 lg:px-8">
			<div class="flex flex-col items-center justify-between gap-4 md:flex-row">
				<p class="text-sm text-center text-gray-500 dark:text-gray-400">
					© {new Date().getFullYear()} Lute's free stack
				</p>
				<div class="flex items-center gap-4">
					<a
						href="https://github.com/lutefd/md-publsiher"
						target="_blank"
						rel="noopener noreferrer"
						class="text-gray-500 hover:text-gray-900 dark:text-gray-400 dark:hover:text-white"
					>
						<span class="sr-only">GitHub</span>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="20"
							height="20"
							viewBox="0 0 24 24"
							fill="none"
							stroke="currentColor"
							stroke-width="2"
							stroke-linecap="round"
							stroke-linejoin="round"
							><path
								d="M15 22v-4a4.8 4.8 0 0 0-1-3.5c3 0 6-2 6-5.5.08-1.25-.27-2.48-1-3.5.28-1.15.28-2.35 0-3.5 0 0-1 0-3 1.5-2.64-.5-5.36-.5-8 0C6 2 5 2 5 2c-.3 1.15-.3 2.35 0 3.5A5.403 5.403 0 0 0 4 9c0 3.5 3 5.5 6 5.5-.39.49-.68 1.05-.85 1.65-.17.6-.22 1.23-.15 1.85v4"
							/><path d="M9 18c-4.51 2-5-2-7-2" /></svg
						>
					</a>
				</div>
			</div>
		</div>
	</footer>
</div>
