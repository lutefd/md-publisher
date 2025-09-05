import { getNotes } from '$lib/notes';
import type { PageLoad } from './$types';

export const load: PageLoad = async () => {
	try {
		const notes = await getNotes();
		return { notes };
	} catch (error) {
		console.error('Error loading notes:', error);
		return { notes: [] };
	}
};
