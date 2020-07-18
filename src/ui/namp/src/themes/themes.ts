import { Theme } from "./theme";
import { darkTheme } from "./dark";
import { lightTheme } from "./light";
import { setCssVar, hexToRgb } from "../util";
import { shadeColor } from "./colorMixer";

export const themes: Record<string, Theme> = {
    'dark': darkTheme,
    'light': lightTheme
}

const intents = ['Primary', 'Success', 'Warning', 'Danger'];

const camelCaseToKebabCase = (str: string) => str.replace(/([a-z0-9])([A-Z])/g, '$1-$2').toLowerCase();

export const applyTheme = (theme: string) => {
    let themeObj = themes[theme];
    const cssColorBlend = (prop: string, amount: number) => hexToRgb(shadeColor(themeObj[prop], amount));
    
    for (let prop of Object.getOwnPropertyNames(themeObj)) {
        setCssVar(`--${camelCaseToKebabCase(prop)}`, hexToRgb(themeObj[prop]));
    }

    for (let intent of intents) {
        setCssVar(`--${intent.toLowerCase()}-hover`, cssColorBlend(`intent${intent}`, -20));
        setCssVar(`--${intent.toLowerCase()}-active`, cssColorBlend(`intent${intent}`, -25));
    }

    setCssVar('--card-shadow', cssColorBlend('backgroundSecondary', -20));
    setCssVar('--dialog-header', cssColorBlend('dialogBackground', 5));
    setCssVar('--cell-background', cssColorBlend('tableBackground', 10));
    setCssVar('--button-background-hover', cssColorBlend('buttonBackground', 10));
    setCssVar('--button-background-active', cssColorBlend('buttonBackground', 20));
}