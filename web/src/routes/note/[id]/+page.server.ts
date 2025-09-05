import { getNote } from '$lib/notes';
import { marked } from 'marked';
import { error } from '@sveltejs/kit';

export async function load({ params }) {
	try {
		const noteId = params.id;
		const note = await getNote(noteId);

		if (!note) {
			throw error(404, 'Note not found');
		}

		const content = marked.parse(note.content);

		return {
			note,
			content
		};
	} catch (err) {
		console.error('Failed to load note:', err);
		throw error(404, 'Note not found');
	}
}
