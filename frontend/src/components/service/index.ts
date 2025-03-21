import { useMountComponent } from '@/hooks/useMountComponent'

import MessageComponent from './Message.vue'
import ModalComponent from './Modal.vue'

// 基础函数
export const useMessageDialog = async (message: string) => {
    return await useMountComponent({ message }).mount(MessageComponent)
}

export const useModal = async (
    title: string,
    content: string,
    confirmButtonText?: string,
    cancelButtonText?: string
): Promise<boolean> => {
    const result = await useMountComponent({
        title,
        content,
        confirmButtonText,
        cancelButtonText
    }).mount<boolean>(ModalComponent)

    return result === true
}
