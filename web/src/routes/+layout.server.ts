import { getNotes } from '$lib/notes';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async () => {
	try {
		const notes = await getNotes();
		return { notes };
	} catch (error) {
		console.error('Error loading notes in layout:', error);
		return { notes: [] };
	}
};
