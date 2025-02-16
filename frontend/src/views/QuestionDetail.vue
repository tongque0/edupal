<template>
    <Layout>
        <div class="container">
            <div class="header">
                <t-button theme="default" variant="text" @click="goBack">
                    <template #icon><t-icon name="arrow-left" /></template>
                    返回
                </t-button>
            </div>
            <t-space direction="vertical" size="large">
                <!-- 问题卡片 -->
                <t-card>
                    <template #title>{{ question.title }}</template>
                    <p  v-html="formatContent(question.content)"></p>
                    <t-button v-if="!showHint" variant="outline" @click="showHint = true" class="mt-4">
                        <template #icon><t-icon name="relation" /></template>
                        显示提示
                    </t-button>
                    <t-alert v-if="showHint" theme="warning" class="mt-4">
                        <template #icon><t-icon name="relation" /></template>
                        {{ question.hint }}
                    </t-alert>
                </t-card>

                <!-- 答案输入卡片 -->
                <t-card>
                    <template #title>你的答案</template>
                    <t-textarea v-model="answer" placeholder="在这里输入你的答案...." :rows="4" class="mb-4" />
                    <t-button theme="primary" @click="handleSubmit">提交答案</t-button>
                </t-card>

                <!-- 判题结果 -->
                <t-card v-if="showExplanation">
                    <template #title>判题结果</template>
                    <t-space>
                        <div>
                            <t-loading v-if="loading" text="加载中..." size="small"></t-loading>
                            <div v-if="!loading">
                                <t-icon name="check-circle-filled" class="text-success" v-if="result.includes('正确')" />
                                <t-icon name="close-circle-filled" class="text-danger" v-if="result.includes('错误')" />
                                <span>{{ result }}</span>
                            </div>
                        </div>
                    </t-space>
                </t-card>
                <!-- 正确答案 -->
                <t-card v-if="showExplanation">
                    <template #title>正确答案</template>
                    <div class="whitespace-pre-line">{{ question.answer }}</div>
                </t-card>
                <!-- 答案解析 -->
                <t-card v-if="showExplanation">
                    <template #title>答案分析</template>
                    <t-loading v-if="loading" text="分析加载中..." size="small"></t-loading>
                    <div v-if="!loading" class="whitespace-pre-line">{{ checkAnalysis }}</div>
                </t-card>

                <!-- AI分析按钮 -->
                <t-button v-if="showExplanation && !aiAnalysis" @click="handleAiAnalysis">
                    <template #icon><t-icon name="help-circle" /></template>
                    获取AI进一步解析
                </t-button>

                <!-- AI解析结果 -->
                <t-card v-if="aiAnalysis">
                    <template #title>AI进一步解析</template>
                    <p>{{ aiAnalysis }}</p>
                </t-card>
            </t-space>
        </div>
    </Layout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import Layout from '@/layout/layout.vue'
import { useRoute,useRouter } from 'vue-router'
import { MessagePlugin } from 'tdesign-vue-next';
import { getQuestions } from '@/api/question'
import { checkAnswer } from '../api/question';
const route = useRoute()
const router = useRouter()  // 添加这行

// 响应式状态
const answer = ref('')
const showHint = ref(false)
const result = ref(null)
const showExplanation = ref(false)
const aiAnalysis = ref(null)
const checkAnalysis = ref()
const question = ref({
    title: '',
    content: '',
    answer: '',
    hint: '',
    explanation: ''
})
// 获取路由参数
onMounted(async () => {
    const id = route.params.id
    console.log('问题ID:', id)
    const response = await getQuestions({ quesid: id })
    if (response && response.questions && response.questions.length > 0) {
        // 获取数组中的第一个问题
        const questionData = response.questions[0]
        console.log('问题数据:', questionData)
        question.value = {
            title: questionData.title || "",
            content: questionData.question || "",
            answer: questionData.answer || "",
            hint: questionData.tip || "",
            explanation: questionData.parse || ""
        }
    }
})

// 方法
const loading = ref(false)

const handleSubmit = async () => {
    if (!answer.value) {
        MessagePlugin.question('请输入你的答案')
        return
    }
    showExplanation.value = true
    loading.value = true
    try {
        const check = await checkAnswer({
            quesid: route.params.id,
            answer: answer.value
        })
        result.value = check.result
        checkAnalysis.value = check.result_msg
    } catch (error) {
        console.error('判题失败:', error)
        MessagePlugin.error('判题失败，请稍后重试')
    } finally {
        loading.value = false
    }
}

const handleAiAnalysis = () => {
    aiAnalysis.value = question.value.explanation
}
const formatContent = (content) => {
    return content.replace(/\\n/g, '<br>');
}


const goBack = () => {
    router.back()
}
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

.t-space {
    width: 800px;
    margin: 0 auto;
}

@media screen and (max-width: 800px) {
    .t-space {
        width: 100%;
    }

}

.mt-4 {
    margin-top: 16px;
}

.mb-4 {
    margin-bottom: 16px;
}

.text-success {
    color: var(--td-success-color);
}
.text-danger {
    color: var(--td-danger-color);
}

.whitespace-pre-line {
    white-space: pre-line;
}
@media (max-width: 768px) {
    .header {
        margin-bottom: 16px;
    }

    /* ...existing mobile styles... */
}
</style>
