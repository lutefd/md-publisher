import typography from '@tailwindcss/typography';

/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			fontFamily: {
				sans: ['InterVariable', 'sans-serif']
			},
			typography: (theme) => ({
				DEFAULT: {
					css: {
						'--tw-prose-body': theme('colors.gray.600'),
						'--tw-prose-headings': theme('colors.gray.900'),
						'--tw-prose-links': theme('colors.blue.600'),
						'--tw-prose-code': theme('colors.gray.800'),
						'--tw-prose-quotes': theme('colors.gray.700'),
						'--tw-prose-pre-bg': theme('colors.gray.900'),
						'--tw-prose-pre-code': theme('colors.gray.200'),

						'--tw-prose-invert-body': theme('colors.gray.400'),
						'--tw-prose-invert-headings': theme('colors.white'),
						'--tw-prose-invert-links': theme('colors.blue.400'),
						'--tw-prose-invert-code': theme('colors.gray.200'),
						'--tw-prose-invert-quotes': theme('colors.gray.300'),
						'--tw-prose-invert-pre-bg': theme('colors.gray.950'),
						'--tw-prose-invert-pre-code': theme('colors.gray.300'),

						img: {
							borderRadius: theme('borderRadius.lg'),
							boxShadow: theme('boxShadow.md'),
							border: `1px solid ${theme('colors.gray.200')}`,
							margin: '1.5rem auto',
							maxWidth: '100%',
							height: 'auto',
							transition: 'all 200ms',
							'&:hover': {
								boxShadow: theme('boxShadow.lg'),
								transform: 'scale(1.01)'
							},
							'.dark &': {
								border: `1px solid ${theme('colors.gray.800')}`
							}
						},
						pre: {
							borderRadius: theme('borderRadius.lg'),
							boxShadow: theme('boxShadow.md'),
							border: `1px solid ${theme('colors.gray.700')}`,
							overflow: 'auto',
							margin: '1.5rem 0',
							'.dark &': {
								border: `1px solid ${theme('colors.gray.800')}`
							}
						},
						'pre code': {
							padding: '1rem 1.25rem',
							display: 'block',
							fontFamily: theme('fontFamily.mono'),
							fontSize: theme('fontSize.sm[0]'),
							background: 'transparent',
							'@media (min-width: 640px)': {
								padding: '1.25rem 1.5rem',
								fontSize: theme('fontSize.base[0]')
							}
						},
						':not(pre) > code': {
							background: theme('colors.gray.100'),
							borderRadius: theme('borderRadius.DEFAULT'),
							padding: '0.2em 0.4em',
							fontSize: '0.9em',
							fontFamily: theme('fontFamily.mono'),
							whiteSpace: 'nowrap',
							'.dark &': {
								background: theme('colors.gray.800'),
								color: theme('colors.gray.200')
							}
						},
						table: {
							width: '100%',
							margin: '1.5rem 0',
							overflow: 'hidden',
							borderRadius: theme('borderRadius.lg'),
							border: `1px solid ${theme('colors.gray.200')}`,
							boxShadow: theme('boxShadow.sm'),
							'.dark &': {
								border: `1px solid ${theme('colors.gray.800')}`
							}
						},
						thead: {
							background: theme('colors.gray.100'),
							'.dark &': {
								background: theme('colors.gray.800')
							}
						},
						th: {
							padding: '0.75rem 1rem',
							textAlign: 'left',
							fontSize: theme('fontSize.sm[0]'),
							fontWeight: theme('fontWeight.medium'),
							borderBottom: `1px solid ${theme('colors.gray.200')}`,
							'.dark &': {
								borderBottom: `1px solid ${theme('colors.gray.700')}`
							}
						},
						'tbody tr': {
							borderBottom: `1px solid ${theme('colors.gray.200')}`,
							'&:last-child': {
								borderBottom: 'none'
							},
							'.dark &': {
								borderBottom: `1px solid ${theme('colors.gray.800')}`
							}
						},
						'tbody tr:nth-child(even)': {
							background: theme('colors.gray.50'),
							'.dark &': {
								background: 'rgba(3, 7, 18, 0.5)'
							}
						},
						td: {
							padding: '0.75rem 1rem',
							fontSize: theme('fontSize.sm[0]')
						},
						blockquote: {
							borderLeftWidth: '4px',
							borderLeftColor: theme('colors.blue.500'),
							background: theme('colors.blue.50'),
							paddingLeft: '1rem',
							paddingTop: '0.5rem',
							paddingBottom: '0.5rem',
							paddingRight: '0.5rem',
							margin: '1.5rem 0',
							borderTopRightRadius: theme('borderRadius.lg'),
							borderBottomRightRadius: theme('borderRadius.lg'),
							'.dark &': {
								borderLeftColor: theme('colors.blue.600'),
								background: 'rgba(30, 58, 138, 0.1)'
							}
						},
						'blockquote p': {
							fontStyle: 'italic',
							color: theme('colors.gray.700'),
							'.dark &': {
								color: theme('colors.gray.300')
							}
						},
						a: {
							color: 'var(--tw-prose-links)',
							fontWeight: theme('fontWeight.medium'),
							transition: 'all 200ms',
							textDecorationThickness: '2px',
							textUnderlineOffset: '2px',
							textDecorationColor: theme('colors.blue.300'),
							'&:hover': {
								color: theme('colors.blue.700'),
								textDecoration: 'underline',
								textDecorationColor: theme('colors.blue.500')
							},
							'.dark &': {
								textDecorationColor: theme('colors.blue.700'),
								'&:hover': {
									color: theme('colors.blue.300'),
									textDecorationColor: theme('colors.blue.500')
								}
							}
						}
					}
				}
			})
		}
	},
	plugins: [typography]
};
