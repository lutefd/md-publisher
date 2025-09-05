import { getNotes } from '$lib/notes';
import type { LayoutLoad } from './$types';

export const load: LayoutLoad = async () => {
	try {
		const notes = await getNotes();
		return { notes };
	} catch (error) {
		console.error('Error loading notes in layout:', error);
		return { notes: [] };
	}
};
