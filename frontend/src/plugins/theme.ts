import { type ThemeDefinition } from 'vuetify'

export const lightTheme: ThemeDefinition = {
    dark: false,
    colors: {
        background: '#FAFAFA',
        surface: '#FFFFFF',
        primary: '#F06292',
        secondary: '#42A5F5',
        error: '#FF5252',
        info: '#2196F3',
        success: '#4CAF50',
        warning: '#FB8C00',
        appbar: '#F06292'
    }
}

export const darkTheme: ThemeDefinition = {
    dark: true,
    colors: {
        background: '#121212',
        surface: '#1E1E1E',
        primary: '#F491B2',
        secondary: '#90CAF9',
        error: '#FF5252',
        info: '#2196F3',
        success: '#4CAF50',
        warning: '#FB8C00',
        appbar: '#212121'
    }
}
