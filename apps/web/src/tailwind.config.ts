import type { Config } from "tailwindcss";

function withOpacity(variableName: string) {
  return ({ opacityValue }: { opacityValue: number }): any => {
    if (opacityValue !== undefined) {
      return `rgba(${variableName}, ${opacityValue})`;
    }
    return `rgb(${variableName})`;
  };
}

export default {
  content: ["./src/**/*.{astro,html,js,jsx,md,mdx,svelte,ts,tsx,vue}"],
  theme: {
    extend: {
      scale: {
        65: ".65",
        70: ".70",
        72.5: ".725",
        75: ".75",
        77.5: ".775",
        80: ".80",
        82.5: ".825",
        85: ".85",
        90: ".90",
        101: "1.01",
        102: "1.02",
        103: "1.03",
        104: "1.04",
      },
      transitionProperty: {
        top: "top",
      },
      zIndex: {
        "-1": "-1",
        500: "500",
        501: "501",
        5000: "5000",
        9999: "9999",
      },
      spacing: {
        128: "32rem",
        144: "36rem",
      },
      width: {
        "dot-base": "0.375rem",
        "dot-side": "0.325rem",
        "dot-side-two": "0.275rem",
        "8xl": "88rem",
        "9xl": "96rem",
        "10xl": "104rem",
      },
      rotate: {
        15: "15deg",
        30: "30deg",
      },
      colors: {
        third: "var(--third-background-color)",
        primary: {
          50: withOpacity("var(--primary-color-50)") as any,
          100: withOpacity("var(--primary-color-100)") as any,
          200: withOpacity("var(--primary-color-200)") as any,
          300: withOpacity("var(--primary-color-300)") as any,
          400: withOpacity("var(--primary-color-400)") as any,
          500: withOpacity("var(--primary-color-500)") as any,
          600: withOpacity("var(--primary-color-600)") as any,
          700: withOpacity("var(--primary-color-700)") as any,
          800: withOpacity("var(--primary-color-800)") as any,
          900: withOpacity("var(--primary-color-900)") as any,
          1000: withOpacity("var(--primary-color-1000)") as any,
          1100: withOpacity("var(--primary-color-1100)") as any,
          1200: withOpacity("var(--primary-color-1200)") as any,
          1300: withOpacity("var(--primary-color-1300)") as any,
          1400: withOpacity("var(--primary-color-1400)") as any,
          DEFAULT: withOpacity("var(--primary-color)") as any,
        },
        secondary: {
          50: withOpacity("var(--secondary-color-50)") as any,
          100: withOpacity("var(--secondary-color-100)") as any,
          200: withOpacity("var(--secondary-color-200)") as any,
          300: withOpacity("var(--secondary-color-300)") as any,
          400: withOpacity("var(--secondary-color-400)") as any,
          500: withOpacity("var(--secondary-color-500)") as any,
          600: withOpacity("var(--secondary-color-600)") as any,
          700: withOpacity("var(--secondary-color-700)") as any,
          800: withOpacity("var(--secondary-color-800)") as any,
          900: withOpacity("var(--secondary-color-900)") as any,
          1000: withOpacity("var(--secondary-color-1000)") as any,
          1100: withOpacity("var(--secondary-color-1100)") as any,
          1200: withOpacity("var(--secondary-color-1200)") as any,
          1300: withOpacity("var(--secondary-color-1300)") as any,
          1400: withOpacity("var(--secondary-color-1400)") as any,
          DEFAULT: withOpacity("var(--secondary-color)") as any,
        },
      },
      borderColor: {
        DEFAULT: "var(--border-color)",
      },
      backgroundImage: {
        "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
        "gradient-conic":
          "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))",
      },
      backdropBlur: {
        xs: "3px",
      },
      fontSize: {
        xxs: "0.625rem",
      },
      backgroundColor: {
        default: "var(--default-background)",
        second: "var(--second-background-color)",
        blured: "var(--blured-background-color)",
      },
      transitionDuration: {
        DEFAULT: "200ms",
      },
      boxShadow: {
        top: "0 -1px 3px 0 var(--shadow-color-start), 0 -4px 8px 3px var(--shadow-color-end)",
        bottom:
          "0 1px 3px 0 var(--shadow-color-start), 0 4px 8px 3px var(--shadow-color-end)",
      },
      height: {
        popup: "calc(100vh - 22rem)",
        "popup-xs": "max(100vh - 30rem, 40vh)",
        "popup-sm": "max(100vh - 25rem, 50vh)",
        "popup-md": "max(100vh - 22rem, 60vh)",
        "popup-lg": "max(100vh - 18rem, 70vh)",
        "popup-xl": "max(100vh - 15rem, 80vh)",
        "popup-2xl": "max(100vh - 12rem, 85vh)",
        "popup-3xl": "max(100vh - 10rem, 90vh)",
        "popup-4xl": "max(100vh - 8rem, 95vh)",
        "sticky-bar": "calc(100vh - 6.5rem)",
        "dot-base": "0.375rem",
        "dot-side": "0.325rem",
        "dot-side-two": "0.275rem",
      },
      top: {
        "sticky-bar": "80px",
      },
      animation: {
        "in-from-top": "inFromTop 0.2s ease-out",
        "in-from-bottom": "inFromBottom 0.2s ease-out",
        "in-from-left": "inFromLeft 0.2s ease-out",
        "in-from-right": "inFromRight 0.2s ease-out",
        "out-to-top": "outToTop 0.2s ease-out",
        "out-to-bottom": "outToBottom 0.2s ease-out",
        "out-to-left": "outToLeft 0.2s ease-out",
        "out-to-right": "outToRight 0.2s ease-out",
        "fade-in": "fadeIn 0.2s ease-out",
        "fade-out": "fadeOut 0.2s ease-out",
        "fade-in-from-top": "fadeIn 0.2s ease-out, inFromTop 0.2s ease-out",
        "fade-in-from-bottom":
          "fadeIn 0.2s ease-out, inFromBottom 0.2s ease-out",
        "fade-in-from-left": "fadeIn 0.2s ease-out, inFromLeft 0.2s ease-out",
        "fade-in-from-right": "fadeIn 0.2s ease-out, inFromRight 0.2s ease-out",
        "fade-out-to-top": "fadeOut 0.2s ease-out, outToTop 0.2s ease-out",
        "fade-out-to-bottom":
          "fadeOut 0.2s ease-out, outToBottom 0.2s ease-out",
        "fade-out-to-left": "fadeOut 0.2s ease-out, outToLeft 0.2s ease-out",
        "fade-out-to-right": "fadeOut 0.2s ease-out, outToRight 0.2s ease-out",
        "slide-down": "inFromTop 0.2s ease-out, fadeIn 0.2s ease-out",
        "slide-up": "outToTop 0.2s ease-out, fadeOut 0.2s ease-out",
      },
      blur: {
        "3xs": "1px",
        xxs: "1.5px",
        xs: "3px",
        "3xl": "30px",
      },
      brightness: {
        10: "0.1",
        15: "0.15",
        20: "0.2",
        25: "0.25",
        30: "0.3",
        35: "0.35",
        40: "0.4",
        45: "0.45",
        55: "0.55",
        60: "0.6",
        65: "0.65",
        70: "0.7",
        75: "0.75",
        80: "0.8",
        85: "0.85",
        90: "0.90",
        95: "0.95",
      },
      keyframes: {
        inFromTop: {
          "0%": {
            transform: "translateY(-100%)",
          },
          "100%": {
            transform: "translateY(0)",
          },
        },
        inFromBottom: {
          "0%": {
            transform: "translateY(100%)",
          },
          "100%": {
            transform: "translateY(0)",
          },
        },
        inFromLeft: {
          "0%": {
            transform: "translateX(-100%)",
          },
          "100%": {
            transform: "translateX(0)",
          },
        },
        inFromRight: {
          "0%": {
            transform: "translateX(100%)",
          },
          "100%": {
            transform: "translateX(0)",
          },
        },
        outToTop: {
          "0%": {
            transform: "translateY(0)",
          },
          "100%": {
            transform: "translateY(-100%)",
          },
        },
        outToBottom: {
          "0%": {
            transform: "translateY(0)",
          },
          "100%": {
            transform: "translateY(100%)",
          },
        },
        outToLeft: {
          "0%": {
            transform: "translateX(0)",
          },
          "100%": {
            transform: "translateX(-100%)",
          },
        },
        outToRight: {
          "0%": {
            transform: "translateX(0)",
          },
          "100%": {
            transform: "translateX(100%)",
          },
        },
        fadeIn: {
          "0%": {
            opacity: "0",
          },
          "100%": {
            opacity: "1",
          },
        },
        fadeOut: {
          "0%": {
            opacity: "1",
          },
          "100%": {
            opacity: "0",
          },
        },
      },
    },
  },
  plugins: [],
} satisfies Config;
