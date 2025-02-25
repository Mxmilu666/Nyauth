<script setup lang="ts">
import stepper from './Stepper.vue'
import nfooter from './Footer.vue'

defineOptions({
    name: 'HomePage'
})
</script>

<template>
    <v-container class="home">
        <v-row justify="center">
            <v-col cols="20">
                <v-row class="fill-height" align="center" justify="center">
                    <v-col cols="12" md="6">
                        <h1 class="font-weight-bold text-h3">Nyauth</h1>
                        <h1 class="font-weight-bold pt-3">喵令在手，秒通万户</h1>
                        <p class="headline pt-2">
                            只需一个 Nyauth 账户即可访问所有已接入 Nyauth 的服务
                        </p>
                        <v-row class="pt-10">
                            <v-btn
                                color="primary"
                                class="ma-2"
                                prepend-icon="mdi-check"
                                @click="$router.push('/login')"
                            >
                                开始使用
                            </v-btn>
                            <v-btn
                                color="secondary"
                                class="ma-2"
                                prepend-icon="mdi-file-document-outline"
                                @click="$router.push('/docs')"
                            >
                                使用文档
                            </v-btn>
                        </v-row>
                    </v-col>
                    <v-col cols="12" md="6">
                        <v-lazy class="d-flex justify-end">
                            <img
                                src="@/assets/banner.png"
                                alt="banner"
                                class="banner-image"
                            />
                        </v-lazy>
                    </v-col>
                </v-row>
            </v-col>
        </v-row>

        <v-row class="pt-10">
            <v-divider />
        </v-row>

        <v-row justify="center">
            <v-col cols="20" class="text-center">
                <h2 class="font-weight-bold text-h4 mt-10">为什么选择 Nyauth?</h2>
                <v-row class="mt-5" justify="center">
                    <v-col cols="12" md="3">
                        <v-card class="card-with-icon" elevation="2">
                            <v-card-title class="font-weight-bold">
                                安全可靠
                            </v-card-title>
                            <v-card-text>
                                Nyauth
                                提供了强大的安全机制，确保用户数据的安全和隐私。我们采用了最新的加密技术和多重身份验证，确保您的账户和数据始终处于安全状态。
                            </v-card-text>
                            <v-icon
                                class="card-icon"
                                size="100"
                                icon="mdi-lock-outline"
                            />
                        </v-card>
                    </v-col>
                    <v-col cols="12" md="3">
                        <v-card class="card-with-icon" elevation="2">
                            <v-card-title class="font-weight-bold">
                                易于集成
                            </v-card-title>
                            <v-card-text>
                                Nyauth
                                具有高度的可扩展性，能够轻松集成到现有系统中。无论是小型应用还是大型企业系统，我们都能提供完美的解决方案。
                            </v-card-text>
                            <v-icon class="card-icon" size="100" icon="mdi-apps" />
                        </v-card>
                    </v-col>
                    <v-col cols="12" md="3">
                        <v-card class="card-with-icon" elevation="2">
                            <v-card-title class="font-weight-bold">
                                高效便捷
                            </v-card-title>
                            <v-card-text>
                                Nyauth
                                提供了简洁的用户界面和高效的操作流程，提升用户体验。我们致力于让每一个用户都能快速上手，享受便捷的服务。
                            </v-card-text>
                            <v-icon class="card-icon" size="100" icon="mdi-speedometer" />
                        </v-card>
                    </v-col>
                    <v-col cols="12" md="3">
                        <v-card class="card-with-icon" elevation="2">
                            <v-card-title class="font-weight-bold">
                                开源驱动
                            </v-card-title>
                            <v-card-text>
                                Nyauth
                                将所有源代码开放，让全世界的开发者们参与到项目中来。我们相信开源社区的力量，能够不断推动项目的发展和进步。
                            </v-card-text>
                            <v-icon class="card-icon" size="100" icon="mdi-earth" />
                        </v-card>
                    </v-col>
                </v-row>
            </v-col>
        </v-row>

        <v-row class="pt-15">
            <v-col cols="12" md="3" justify="center">
                <h2 class="font-weight-bold text-h4 mt-10">只需四步即可接入 Nyauth</h2>
            </v-col>
            <v-col cols="12" md="9">
                <stepper />
            </v-col>
            <v-divider class="pt-10" />
        </v-row>
    </v-container>
    <nfooter />
</template>

<style scoped>
.home {
    .banner-image {
        width: 66%;
        height: auto;
    }

    .card-icon {
        position: absolute;
        bottom: -36px;
        right: -36px;
        opacity: 0.1;
        transform: scale(2);
    }

    .card-with-icon {
        padding: 20px;
        min-height: 220px;
    }
}
</style>
