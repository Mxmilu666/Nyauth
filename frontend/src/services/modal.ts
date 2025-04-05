import { useModal } from '@/components/service/index'
import type { ModalButton } from '@/components/service/index'

export const modal = {
    show<T = boolean>(options: {
        title: string
        content: string
        buttons?: ModalButton[]
        prependIcon?: string
    }): Promise<T> {
        return useModal<T>(
            options.title,
            options.content,
            {
                buttons: options.buttons,
                prependIcon: options.prependIcon
            }
        )
    },

    error<T = boolean>(options: { 
        title?: string; 
        content: string;
        buttons?: ModalButton[];
        prependIcon?: string
    }): Promise<T> {
        const { title = '错误', content } = options
        return useModal<T>(title, content, {
            buttons: options.buttons || [
                { text: '确定', color: 'primary', variant: 'text', value: true }
            ],
            prependIcon: options.prependIcon || 'mdi-close-circle'
        })
    },

    success<T = boolean>(options: { 
        title?: string; 
        content: string;
        buttons?: ModalButton[];
        prependIcon?: string
    }): Promise<T> {
        const { title = '成功', content } = options
        return useModal<T>(title, content, {
            buttons: options.buttons || [
                { text: '确定', color: 'success', variant: 'text', value: true }
            ],
            prependIcon: options.prependIcon || 'mdi-check-circle'
        })
    },

    confirm<T = boolean>(options: { 
        title?: string; 
        content: string;
        buttons?: ModalButton[];
        prependIcon?: string
    }): Promise<T> {
        const { title = '确认', content } = options
        return useModal<T>(title, content, {
            buttons: options.buttons,
            prependIcon: options.prependIcon || 'mdi-help-circle'
        })
    }
}