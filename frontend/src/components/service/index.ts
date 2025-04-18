import { useMountComponent } from '@/hooks/useMountComponent'

import MessageComponent from './Message.vue'
import ModalComponent from './Modal.vue'

// 定义消息类型
export type MessageType = 'info' | 'success' | 'error' | 'warning'

// 定义按钮类型接口
export interface ModalButton {
    text: string;
    color?: string;
    variant?: string;
    value: any;
}

// 基础函数
export const useMessageDialog = async (message: string, type: MessageType = 'info') => {
    return await useMountComponent({ message, type }).mount(MessageComponent)
}

export const useModal = async <T = boolean>(
    title: string,
    content: string,
    options?: {
        buttons?: ModalButton[];
        prependIcon?: string;
    }
): Promise<T> => {
    const result = await useMountComponent({
        title,
        content,
        prependIcon: options?.prependIcon || 'mdi-update',
        buttons: options?.buttons || [
            { text: '取消', color: 'error', variant: 'text', value: false },
            { text: '确认', color: 'primary', variant: 'text', value: true }
        ]
    }).mount<T>(ModalComponent)

    return result as T
}