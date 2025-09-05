import { getNotes } from '$lib/notes';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async () => {
	try {
		const notes = await getNotes();
		return { notes };
	} catch (error) {
		console.error('Error loading notes:', error);
		return { notes: [] };
	}
};
