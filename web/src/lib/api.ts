export interface Note {
	id: string;
	content: string;
	metadata: {
		title?: string;
		description?: string;
		tags?: string[];
		updated?: string;
		[key: string]: unknown;
	};
}

let API_URL = '/api';

if (typeof window === 'undefined') {
	API_URL = 'http://publisher-api:8080';
} else {
	if (import.meta.env.DEV) {
		API_URL = 'http://publisher-api:8080';
	}
}

/**
 * Fetch all published notes
 */
export async function getAllNotes(): Promise<Note[]> {
	const response = await fetch(`${API_URL}/notes`);

	if (!response.ok) {
		throw new Error(`Failed to fetch notes: ${response.statusText}`);
	}

	return response.json();
}

/**
 * Fetch a specific note by ID
 */
export async function getNoteById(id: string): Promise<Note> {
	const response = await fetch(`${API_URL}/note/${id}`);

	if (!response.ok) {
		throw new Error(`Failed to fetch note: ${response.statusText}`);
	}

	return response.json();
}

/**
 * Process note content to separate frontmatter and body
 * This is useful when you want to display just the content without frontmatter
 */
export function processNoteContent(note: Note): {
	frontmatter: Record<string, unknown>;
	content: string;
} {
	return {
		frontmatter: note.metadata,
		content: note.content
	};
}
