<template>
    <Layout>
        <div class="home">
            <section class="hero-section">
                <h1 class="title">鞍山市高新区实验学校 - AI驱动的智能学习平台</h1>
                <div class="feature-list" v-if="!isMobile">
                    <p class="feature-item">
                        <t-icon name="root-list" class="feature-icon" />
                        智能出题：根据你的学习进度定制题目
                    </p>
                    <p class="feature-item">
                        <t-icon name="check-circle-filled" class="feature-icon" />
                        即时判题：获得详细的答案解析和改进建议
                    </p>
                </div>
            </section>

            <section class="quick-start-section">
                <h2 class="section-title">快速开始</h2>
                <div class="form-grid">
                    <t-select v-model="subject" creatable filterable :options="optionsMap.subject"
                        @create="(val) => createOptions(val, 'subject')" placeholder="选择科目">
                    </t-select>

                    <t-select v-model="grade" creatable filterable :options="optionsMap.grade"
                        @create="(val) => createOptions(val, 'grade')" placeholder="选择年级">
                    </t-select>

                    <t-select v-model="difficulty" creatable filterable :options="optionsMap.difficulty"
                        @create="(val) => createOptions(val, 'difficulty')" placeholder="选择难度">
                    </t-select>
                </div>

                <div class="question-types">
                    <h3 class="subsection-title">选择题型和数量</h3>
                    <div class="types-grid">
                        <div v-for="type in questionTypes" :key="type" class="type-card">
                            <t-checkbox v-model="selectedTypes[type]">{{ type }}</t-checkbox>
                            <div v-if="selectedTypes[type]" class="counter">
                                <t-input-number v-model="typeCounts[type]" :min="1" />
                            </div>
                        </div>
                    </div>
                </div>

                <div class="action-row">
                    <t-input v-model="topic" placeholder="输入你想练习的知识点" />
                    <t-button theme="primary" @click="handleStart">
                        开始练习
                        <template #suffix><t-icon name="chevron-right" /></template>
                    </t-button>
                </div>
            </section>
        </div>
    </Layout>
</template>

<script setup>
import { ref ,onMounted} from 'vue'
import Layout from '@/layout/layout.vue'
import { MessagePlugin } from 'tdesign-vue-next';
import { useRouter } from 'vue-router'
import { genQuestions } from '../api/question'
const router = useRouter()
const subjects = ['数学', '物理', '化学', '生物', '语文', '英语']
const grades = ['初一', '初二', '初三', '高一', '高二', '高三']
const difficulties = ['简单', '中等', '困难']
const questionTypes = ['选择题','计算题', '简答题']

const optionsMap = ref({
    subject: subjects.map(item => ({ label: item, value: item })),
    grade: grades.map(item => ({ label: item, value: item })),
    difficulty: difficulties.map(item => ({ label: item, value: item })),
})

const createOptions = (val, type) => {
    optionsMap.value[type].push({
        label: val,
        value: val
    })
}

const subject = ref('')
const grade = ref('')
const difficulty = ref('')
const topic = ref('')
const selectedTypes = ref({})
const typeCounts = ref({
    '选择题': 1,
    '计算题': 1,
    '简答题': 1,
})


const handleStart = async () => {
    if (!topic.value) {
        MessagePlugin.question('请输入你想练习的知识点')
        return
    }
    if (!subject.value || !grade.value || !difficulty.value) {
        MessagePlugin.question('请选择科目、年级和难度')
        return
    }
    if (!Object.values(selectedTypes.value).some(Boolean)) {
        MessagePlugin.question('请选择至少一种题型')
        return
    }

    const requestData = {
        subject: subject.value,
        class: grade.value,
        difficulty: difficulty.value,
        knowledge: topic.value,
        choice_num: selectedTypes.value['选择题'] ? typeCounts.value['选择题'] : 0,
        calc_num: selectedTypes.value['计算题'] ? typeCounts.value['计算题'] : 0,
        short_ans_num: selectedTypes.value['简答题'] ? typeCounts.value['简答题'] : 0,
    }

    try {
        const response = await genQuestions(requestData)

        router.push(`/preview/${response.genid}`)

    } catch (error) {
        MessagePlugin.error('生成题目失败，请稍后重试')
    }
}

const isMobile = ref(false);
const checkScreenWidth = () => {
    isMobile.value = window.innerWidth < 768;
};
onMounted(() => {
    checkScreenWidth();
    window.addEventListener('resize', checkScreenWidth);
});
</script>
<style scoped>
.home {
    padding: 16px;
    margin: 0 auto;
    width: 90%;
    max-width: 1600px;
    display: flex;
    /* 添加 flex 布局 */
    flex-direction: column;
    align-items: center;
    /* 水平居中 */
}

.hero-section {
    background: linear-gradient(135deg, #0052D9 0%, #0066ff 100%);
    /* 添加渐变背景 */
    color: white;
    padding: 1.5rem 2rem;
    /* 调整内边距 */
    border-radius: 0.5rem;
    text-align: center;
    margin: 0.5rem auto;
    width: 100%;
    max-width: 1200px;
    /* 限制最大宽度 */
    box-shadow: 0 4px 12px rgba(0, 82, 217, 0.15);
    /* 优化阴影效果 */
}

.title {
    font-size: 2rem;
    font-weight: bold;
    margin-bottom: 1rem;
}

.feature-list {
    display: flex;
    flex-direction: row;
    justify-content: center;
    /* 确保特性列表居中 */
    gap: 3rem;
    /* 增加间距 */
    margin: 0 auto;
    /* 居中对齐 */
    max-width: 800px;
    /* 限制最大宽度 */
}

.feature-item {
    font-size: 1rem;
    display: flex;
    align-items: center;
    gap: 0.5rem;
}

.feature-icon {
    font-size: 1.25rem;
}

@media (max-width: 768px) {
    .home {
        width: 95%;
        /* 移动端略微增加宽度占比 */
        padding: 12px;
    }

    .hero-section {
        padding: 1rem;
    }

    .feature-list {
        flex-direction: column;
        /* 移动端改回垂直排列 */
        gap: 0.5rem;
    }

    .title {
        font-size: 1.5rem;
    }
}

.quick-start-section {
    width: 100%;
    /* 确保宽度为100% */
    max-width: 1200px;
    /* 限制最大宽度 */
    background: white;
    border-radius: 0.5rem;
    padding: 2rem;
    margin: 2rem auto;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.section-title {
    font-size: 1.5rem;
    font-weight: bold;
    margin-bottom: 1.5rem;
    color: #1a1a1a;
}

.form-grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 1rem;
    margin-bottom: 2rem;
}

.subsection-title {
    font-size: 1.2rem;
    font-weight: 600;
    margin-bottom: 1rem;
    color: #1a1a1a;
}

.types-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    gap: 1rem;
    margin-bottom: 2rem;
}

.type-card {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1rem;
    border: 1px solid #e5e7eb;
    border-radius: 0.5rem;
}

.counter {
    display: flex;
    align-items: center;
    gap: 0.5rem;
}

.action-row {
    display: flex;
    gap: 1rem;
    align-items: center;
}

@media (max-width: 768px) {
  .home {
    padding: 0 16px; /* 增加左右边距 */
    width: 85%;
    box-sizing: border-box; /* 防止溢出 */
  }

  .hero-section {
    margin: 0.5rem -16px; /* 扩展背景到屏幕边缘 */
    width: calc(100% + 32px);
    padding: 1.5rem 16px;
    border-radius: 0; /* 移除圆角更符合移动端风格 */
  }

  .quick-start-section {
    padding: 1.5rem;
    margin: 1rem -16px; /* 扩展内容区域 */
    width: calc(100% + 32px);
    border-radius: 0;
    box-shadow: none; /* 移除阴影保持简洁 */
    border-top: 1px solid #eee;
    border-bottom: 1px solid #eee;
  }

  .form-grid {
    gap: 12px; /* 增加表单项间距 */
  }

  .types-grid {
    gap: 8px;
  }

  .type-card {
    padding: 12px 16px;
    border-radius: 8px; /* 调整圆角大小 */
  }

  .action-row {
    gap: 12px;
  }
  .form-grid {
        display: flex;
        flex-direction: column;
        gap: 12px;
    }
  /* 优化输入元素尺寸 */
  .t-input, .t-select, .t-button {
    font-size: 16px !important; /* 增大字体 */
    height: 48px !important; /* 增大点击区域 */
  }

  .t-checkbox {
    font-size: 15px;
  }

  /* 调整标题尺寸 */
  .title {
    font-size: 1.25rem;
    line-height: 1.4;
    padding: 0 10px;
  }

  .section-title {
    font-size: 1.3rem;
    margin-bottom: 1.2rem;
  }

  .subsection-title {
    font-size: 1.1rem;
  }

  /* 优化计数器输入框 */
  .t-input-number {
    width: 100px;
  }

  .t-input-number .t-input {
    height: 40px !important;
  }
}

/* 全局移动端优化 */
@media (max-width: 768px) {
  .t-select__wrap, .t-input__wrap {
    width: 100% !important; /* 强制表单元素占满宽度 */
  }

  .t-button {
    width: 100%; /* 按钮占满宽度 */
    justify-content: center;
  }
}
</style>
