<template>
    <Layout>
        <div class="container">
            <div class="header">
                <t-button theme="default" variant="text" @click="goBack">
                    <template #icon><t-icon name="arrow-left" /></template>
                    返回
                </t-button>
                <h3 class="page-title">预览题目</h3>
            </div>
            <t-loading :loading="loading">
                <template v-if="questions.length > 0">
                    <t-space direction="vertical" size="large">
                        <t-card
                            v-for="question in questions"
                            :key="question.id"
                            class="question-card"
                            @click="goToQuestion(question.id)"
                            hoverable
                        >
                            <template #title>
                                <div class="card-title">
                                    {{ question.title }}
                                    <t-icon name="chevron-right" />
                                </div>
                            </template>
                            <p>{{ question.content }}</p>
                            <div class="mt-2">
                                <t-tag v-for="tag in question.tags" :key="tag" class="mr-2">
                                    {{ tag }}
                                </t-tag>
                            </div>
                        </t-card>
                    </t-space>
                </template>
                <template v-else>
                    <div class="text-center py-8">
                        <t-loading :loading="true" text="正在生成题目，请稍候..." />
                    </div>
                </template>
            </t-loading>
        </div>
    </Layout>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import Layout from '@/layout/layout.vue'
import { useRoute, useRouter } from 'vue-router'
import { MessagePlugin } from 'tdesign-vue-next'
import { getQuestions } from '@/api/question'

const route = useRoute()
const router = useRouter()
const questions = ref([])
const loading = ref(true)
const pollingInterval = ref(null)

const fetchQuestions = async () => {
    try {
        const id = route.params.id
        const response = await getQuestions({ genid: id })

        if (response && response.questions) {
            // 如果 response 中有 status 或类似字段表示生成完成
            if (!response.has_more ) {
                clearInterval(pollingInterval.value)
                loading.value = false
            }

            questions.value = response.questions.map(question => ({
                id: question.ques_id,
                title: question.title,
                content: question.question,
                author: question.owner_id,
                tags: question.tags || [],
                createdAt: question.createdAt || '',
                difficulty: question.difficulty || ''
            }))
        }
    } catch (error) {
        console.error('获取题目失败:', error)
        MessagePlugin.error('获取题目失败，请稍后重试')
        clearInterval(pollingInterval.value)
        loading.value = false
    }
}
const goToQuestion = (id) => {
    router.push(`/question/${id}`)
}
const goBack = () => {
    router.back()
}
onMounted(() => {
    // 立即执行一次
    fetchQuestions()

    // 设置轮询间隔（每3秒查询一次）
    pollingInterval.value = setInterval(fetchQuestions, 500)

    // 设置超时时间（比如2分钟后停止轮询）
    setTimeout(() => {
        if (pollingInterval.value) {
            clearInterval(pollingInterval.value)
            if (questions.value.length === 0) {
                MessagePlugin.warning('生成题目超时，请重试')
            }
            loading.value = false
        }
    }, 120000) // 2分钟超时
})

// 组件卸载时清理轮询
onUnmounted(() => {
    if (pollingInterval.value) {
        clearInterval(pollingInterval.value)
    }
})
</script>

<style scoped>
.header {
    display: flex;
    align-items: center;
    margin-bottom: 24px;
    gap: 16px;
}

.page-title {
    margin: 0;
    font-size: 1.25rem;
    font-weight: 500;
}

.container {
    max-width: 800px;
    margin: 0 auto;
    padding: 24px;
}

/* 添加问题卡片样式 */
.question-card {
    cursor: pointer;
    transition: transform 0.2s ease;
}

.question-card:hover {
    transform: translateY(-2px);
}

.card-title {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

/* 保持原有样式 */
.mt-2 {
    margin-top: 8px;
}

.mr-2 {
    margin-right: 8px;
}

.text-center {
    text-align: center;
}

.py-8 {
    padding-top: 32px;
    padding-bottom: 32px;
}

/* 响应式调整 */
@media (max-width: 768px) {
    .header {
        margin-bottom: 16px;
    }
    .container {
        padding: 16px;
    }

    .question-card {
        margin-bottom: 12px;
    }
}
</style>
