<script setup lang="ts">
import { ref } from 'vue'

defineOptions({
    name: 'StepperComponent'
})

const steps = [
    {
        title: '创建账户',
        description: '注册一个新的 Nyauth 账户，填写基本信息并验证邮箱。'
    },
    {
        title: '配置应用',
        description: '在 Nyauth 控制台中添加并配置你的应用，获取客户端 ID 和密钥。'
    },
    {
        title: '集成 Nyauth',
        description: '将 Nyauth 集成到你的应用中，使用 OAuth 2.0 完成身份验证。'
    },
    { title: '测试和发布', description: '测试集成效果，确保一切正常后发布你的应用。' }
]

const step_n = ref(1)
const next = () => {
    if (step_n.value < steps.length) step_n.value += 1
}

const prev = () => {
    if (step_n.value > 1) step_n.value -= 1
}

</script>

<template>
    <v-stepper :model-value="step_n">
        <v-stepper-header>
            <v-stepper-item
                v-for="(step, i) in steps"
                :key="i"
                :title="step.title"
                :value="i + 1"
                :complete="step_n > i + 1"
                editable
            />
        </v-stepper-header>

        <v-stepper-window>
            <v-stepper-window-item
                v-for="(step, i) in steps"
                :key="`${i}-content`"
                :value="i + 1"
            >
                {{ step.description }}
            </v-stepper-window-item>
        </v-stepper-window>

        <v-stepper-actions @click:next="next" @click:prev="prev" />
    </v-stepper>
</template>
