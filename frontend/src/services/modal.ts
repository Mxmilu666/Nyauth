import { useModal } from '@/components/service/index'

export const modal = {
    show(options: {
        title: string
        content: string
        confirmButtonText?: string
        cancelButtonText?: string
    }): Promise<boolean> {
        return useModal(
            options.title,
            options.content,
            options.confirmButtonText,
            options.cancelButtonText
        )
    },

    error(options: { title?: string; content: string }): Promise<boolean> {
        const { title = '错误', content } = options
        return useModal(title, content, '确定', '取消')
    },

    success(options: { title?: string; content: string }): Promise<boolean> {
        const { title = '成功', content } = options
        return useModal(title, content, '确定', '取消')
    },

    confirm(options: { title?: string; content: string }): Promise<boolean> {
        const { title = '确认', content } = options
        return useModal(title, content)
    }
}
