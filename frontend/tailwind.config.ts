import type { Config } from 'tailwindcss';

const config: Config = {
  content: ['./src/**/*.{js,ts,jsx,tsx}'],
  darkMode: 'class',
  theme: {
    extend: {
      width: {
        90: '360px',
        108: '32rem',
        sidebar: '280px',
        base_open: 'calc(100vw - 560px)',
        base_close: 'calc(100vw - 380px)',
        taskbar: '720px',
        taskbar_md: '90%',
        bottomBar: '100px',
      },
      height: {
        90: '360px',
        108: '32rem',
        navbar: '64px',
        base: 'calc(100vh - 64px)',
        taskbar: '48px',
        base_md: 'calc(100vh - 64px - 48px)',
      },
      minHeight: {
        base_md: 'calc(100vh - 64px - 48px)',
      },
      spacing: {
        navbar: '64px',
        base_padding: '24px',
        bottomBar: '100px',
        base_md: 'calc(100vh - 64px - 48px)',
      },
      fontFamily: {
        primary: ['var(--inter-font)'],
      },
      fontSize: {
        xxs: '0.5rem',
      },
      colors: {
        primary_text: '#478EE1',
        dark_primary_gradient_start: '#633267',
        dark_primary_gradient_end: '#5b406b',
        dark_secondary_gradient_start: '#be76bf',
        dark_secondary_gradient_end: '#607ee7',
        primary_btn: '#9ca3af',
        dark_primary_btn: '#9275b9ba',
        primary_comp: '#478eeb18',
        dark_primary_comp: '#c578bf1b',
        primary_comp_hover: '#478eeb38',
        dark_primary_comp_hover: '#c578bf36',
        primary_comp_active: '#478eeb86',
        dark_primary_comp_active: '#c578bf5d',
        primary_danger: '#ea333e',
        primary_black: '#2e2c2c',
        heart_filled: '#fe251baa',
        priority_high: '#fbbebe',
        priority_mid: '#fbf9be',
        priority_low: '#bffbbe',
      },
      backgroundColor: {
        backdrop: '#0000003f',
        navbar: '#ffffff',
        main: '#e5e7eb',
        sidebar: '#ffffff',
      },
      animation: {
        fade_third: 'fade 0.3s ease-in-out',
        fade_third_delay: 'fade 0.3s ease-in-out 0.5s',
        fade_half: 'fade 0.5s ease-in-out',
        fade_1: 'fade 1s ease-in-out',
        fade_2: 'fade 2s ease-in-out',
        shrink: 'shrink 0.1s ease-in-out 0.4s forwards',
      },
      keyframes: {
        shrink: {
          '0%': { scale: '100%' },
          '100%': { scale: '0%' },
        },
        fade: {
          '0%': { opacity: '0' },
          '100%': { opacity: '1' },
        },
      },
    },
  },
  plugins: [],
};
export default config;
