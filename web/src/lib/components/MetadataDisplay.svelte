<script lang="ts">
	const { metadata = {} } = $props<{ metadata: Record<string, any> }>();

	function isList(value: any): boolean {
		return Array.isArray(value);
	}

	function isSimpleValue(value: any): boolean {
		return typeof value === 'string' || typeof value === 'number' || typeof value === 'boolean';
	}

	function isObject(value: any): boolean {
		return typeof value === 'object' && value !== null && !Array.isArray(value);
	}

	const filteredMetadata = $derived(
		(() => {
			const result = { ...metadata };
			delete result.title;
			delete result.description;
			delete result.tags;
			return result;
		})()
	);

	const hasMetadata = $derived((() => Object.keys(filteredMetadata).length > 0)());
</script>

{#if hasMetadata}
	<div
		class="p-4 mt-6 border border-gray-200 rounded-lg bg-gray-50 dark:border-gray-800 dark:bg-gray-900/50"
	>
		<h3 class="mb-3 text-sm font-medium text-gray-900 dark:text-gray-100">Metadata</h3>

		<div class="space-y-3">
			{#each Object.entries(filteredMetadata) as [key, value]}
				<div>
					<span class="text-xs font-semibold text-gray-700 dark:text-gray-300">{key}:</span>

					{#if isSimpleValue(value)}
						<span class="ml-2 text-sm text-gray-600 dark:text-gray-400">{value}</span>
					{:else if Array.isArray(value)}
						<div class="flex flex-wrap gap-1 mt-1">
							{#each value as item}
								{#if isSimpleValue(item)}
									<span
										class="inline-flex items-center px-2 py-1 text-xs font-medium text-gray-600 bg-gray-100 rounded-md dark:bg-gray-800 dark:text-gray-300"
									>
										{item}
									</span>
								{:else}
									<span
										class="inline-flex items-center px-2 py-1 text-xs font-medium text-gray-600 bg-gray-100 rounded-md dark:bg-gray-800 dark:text-gray-300"
									>
										[Complex value]
									</span>
								{/if}
							{/each}
						</div>
					{:else if isObject(value)}
						<details class="mt-1">
							<summary class="text-xs text-blue-600 cursor-pointer dark:text-blue-400">
								View details
							</summary>
							<pre
								class="p-2 mt-2 overflow-auto text-xs bg-gray-100 rounded max-h-40 dark:bg-gray-800">
								{JSON.stringify(value, null, 2)}
							</pre>
						</details>
					{:else}
						<span class="ml-2 text-sm italic text-gray-500 dark:text-gray-500">
							[Complex value]
						</span>
					{/if}
				</div>
			{/each}
		</div>
	</div>
{/if}
