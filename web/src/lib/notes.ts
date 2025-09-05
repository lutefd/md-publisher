import { getAllNotes, getNoteById, type Note } from './api';

/**
 * Get all notes from the API
 */
export async function getNotes(): Promise<Note[]> {
	try {
		const notes = await getAllNotes();
		return notes;
	} catch (error) {
		console.error('Failed to fetch notes:', error);
		return [];
	}
}

/**
 * Get a note by ID
 */
export async function getNote(id: string): Promise<Note | undefined> {
	try {
		return await getNoteById(id);
	} catch (error) {
		console.error(`Failed to fetch note ${id}:`, error);
		return undefined;
	}
}

/**
 * Get notes filtered by tag
 */
export async function getNotesByTag(tag: string): Promise<Note[]> {
	const notes = await getNotes();
	return notes.filter(
		(note) =>
			note.metadata.tags && Array.isArray(note.metadata.tags) && note.metadata.tags.includes(tag)
	);
}

/**
 * Get all unique tags from notes
 */
export async function getAllTags(): Promise<string[]> {
	const notes = await getNotes();
	const tagsSet = new Set<string>();

	notes.forEach((note) => {
		if (note.metadata.tags && Array.isArray(note.metadata.tags)) {
			note.metadata.tags.forEach((tag) => tagsSet.add(tag));
		}
	});

	return Array.from(tagsSet);
}

/**
 * Search notes by query
 */
export async function searchNotes(query: string): Promise<Note[]> {
	const notes = await getNotes();
	const lowerQuery = query.toLowerCase();

	return notes.filter((note) => {
		if (note.metadata.title && note.metadata.title.toLowerCase().includes(lowerQuery)) {
			return true;
		}
		if (note.metadata.description && note.metadata.description.toLowerCase().includes(lowerQuery)) {
			return true;
		}
		if (note.content.toLowerCase().includes(lowerQuery)) {
			return true;
		}
		if (
			note.metadata.tags &&
			Array.isArray(note.metadata.tags) &&
			note.metadata.tags.some((tag) => tag.toLowerCase().includes(lowerQuery))
		) {
			return true;
		}
		if (
			note.metadata.tags &&
			Array.isArray(note.metadata.tags) &&
			note.metadata.tags.some((tag) => tag.toLowerCase().includes(lowerQuery))
		) {
			return true;
		}

		return false;
	});
}
