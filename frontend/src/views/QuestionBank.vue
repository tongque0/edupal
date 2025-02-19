<template>
    <Layout>
        <div class="question-bank-container">
            <!-- 侧边栏分类 -->
            <aside class="sidebar">
                <div class="category-list">
                    <t-menu theme="light" :value="selectedCategory" @change="handleCategoryChange">
                        <t-menu-item v-for="category in categories" :key="category.id" :value="category.id">
                            <template #icon>{{ category.icon }}</template>
                            {{ category.name }}
                        </t-menu-item>
                    </t-menu>
                </div>
            </aside>

            <!-- 主内容区 -->
            <main class="main-content">
                <!-- 搜索栏和视图切换 -->
                <div class="top-bar">
                    <t-input v-model="searchTerm" placeholder="搜索题库..." class="search-input">
                        <template #suffix><t-icon name="search" /></template>
                    </t-input>
                    <div class="view-mode-toggle">
                        <t-button :variant="viewMode === 'questions' ? 'secondary' : 'ghost'"
                            @click="viewMode = 'questions'">
                            <div> <t-icon name="list" /> 题目</div>
                        </t-button>
                        <t-button :variant="viewMode === 'banks' ? 'secondary' : 'ghost'" @click="viewMode = 'banks'">
                            <div> <t-icon name="catalog" /> 题库</div>
                        </t-button>
                        <t-button @click="createQuestionBank">
                            <div> <t-icon name="add" /> 创建题库</div>
                        </t-button>
                    </div>
                </div>

                <!-- 移动端分类标签 -->
                <div class="mobile-categories">
                    <div class="categories-scroll">
                        <t-tag v-for="category in categories" :key="category.id"
                            :theme="selectedCategory === category.id ? 'primary' : 'default'" variant="light"
                            class="category-tag" @click="handleCategoryChange(category.id)">
                            {{ category.icon }} {{ category.name }}
                        </t-tag>
                    </div>
                </div>

                <!-- 内容列表 -->
                <div class="content-list">
                    <template v-if="viewMode === 'questions'">
                        <t-card v-for="question in questions" :key="question.id" :hover="true" class="content-card">
                            <div class="content-main">
                                <h3 class="content-title" @click="goToQuestionDetail(question.id)">{{ question.title }}</h3>
                                <p class="content-description" v-html="formatContent(question.content)"></p>
                                <div class="content-meta">
                                    <span class="author">{{ question.author }}</span>
                                    <span class="views">
                                        <t-icon name="browse" />{{ question.views }}
                                    </span>
                                    <span class="likes">
                                        <t-icon name="thumb-up" />{{ question.likes }}
                                    </span>
                                    <span class="comments">
                                        <t-icon name="message-square" />{{ question.comments }}
                                    </span>
                                    <span class="created-at">
                                        <t-icon name="clock" />{{ question.createdAt }}
                                    </span>
                                    <t-badge v-for="tag in question.tags" :key="tag" variant="outline">{{ tag
                                        }}</t-badge>
                                </div>
                            </div>
                        </t-card>
                    </template>
                    <template v-else>
                        <t-card v-for="bank in questionBanks" :key="bank.id" :hover="true" class="content-card">
                            <div class="content-main">
                                <h3 class="content-title" @click="goToQuesBankDetail(bank.id)">{{ bank.title }}</h3>
                                <p class="content-description">{{ bank.description }}</p>
                                <div class="content-meta">
                                    <span class="author">{{ bank.author }}</span>
                                    <span class="views">
                                        <t-icon name="browse" />{{ bank.views }}
                                    </span>
                                    <span class="likes">
                                        <t-icon name="thumb-up" />{{ bank.likes }}
                                    </span>
                                    <span class="count">
                                        <t-icon name="book-open" />{{ bank.questionCount }}题
                                    </span>
                                    <span class="created-at">
                                        <t-icon name="clock" />{{ bank.createdAt }}
                                    </span>
                                    <t-badge v-for="tag in bank.tags" :key="tag" variant="outline">{{ tag }}</t-badge>
                                </div>
                            </div>
                        </t-card>
                    </template>
                </div>
            </main>
        </div>
    </Layout>
</template>

<script setup>
import { ref ,computed,watch, onMounted} from 'vue'
import { useRouter } from 'vue-router'
import Layout from '@/layout/layout.vue'
import { getQuestions } from '@/api/question'

const router = useRouter()
const searchTerm = ref('')
const selectedCategory = ref('')
const viewMode = ref('questions')
const questionBanks = ref([])
const questions = ref([])
const categories = [
    { id: 'follow', name: '关注', icon: '👀' },
    { id: 'all', name: '综合', icon: '🔥' },
    { id: 'math', name: '数学', icon: '📐' },
    { id: 'physics', name: '物理', icon: '⚡' },
    { id: 'chemistry', name: '化学', icon: '🧪' },
    { id: 'biology', name: '生物', icon: '🧬' },
    { id: 'chinese', name: '语文', icon: '📚' },
    { id: 'english', name: '英语', icon: '🌍' }
]

//构造查询条件，在条件变化时触发查询
const queryQuesConditions = computed(() => {
    const selectedCategoryObj = categories.find(category => category.id === selectedCategory.value)
    return {
        subject: selectedCategoryObj && selectedCategoryObj.name !== '综合' ? selectedCategoryObj.name : ''
    }
})

onMounted(() => {
    selectedCategory.value = 'all'
})

watch(queryQuesConditions, async (queryQuesConditions) => {
    questions.value = []
    const response = await getQuestions(queryQuesConditions)
    if (response && response.questions) {
        questions.value = response.questions.map(question => ({
            id: question.ques_id,
            title: question.title,
            content: question.question,
            author: question.owner_id,
            views: question.views || 0,
            likes: question.likes || 0,
            comments: question.comments || 0,
            tags: question.tags || [],
            createdAt: question.createdAt || '',
            difficulty: question.difficulty || ''
        }))
    }
})


// const questionBanks = [
//     {
//         id: '1',
//         title: '高中物理力学综合题库',
//         description: '包含牛顿运动定律、能量守恒等经典力学题目，难度适中，适合高一学生。',
//         author: '物理老王',
//         views: 1200,
//         likes: 45,
//         tags: ['物理', '力学', '高中'],
//         questionCount: 50,
//         createdAt: '2024-02-10'
//     },
//     {
//         id: '2',
//         title: '中考数学重点复习题库',
//         description: '针对中考常见题型，包含详细解析和答题技巧。',
//         author: '数学老张',
//         views: 890,
//         likes: 32,
//         tags: ['数学', '中考', '初三'],
//         questionCount: 100,
//         createdAt: '2024-02-12'
//     }
// ]

const handleCategoryChange = (categoryId) => {
    selectedCategory.value = categoryId
}

const formatContent = (content) => {
    return content.replace(/\\n/g, '<br>');
}
const goToQuestionDetail = (id) => {
    router.push(`/question/${id}`)
}
const goToQuesBankDetail = (id) => {
    router.push(`/question-bank/${id}`)
}

const createQuestionBank = () => {
    // 创建题库的逻辑
}
</script>

<style scoped>
.question-bank-container {
    display: flex;
    min-height: 100vh;
    background-color: #f5f5f5;
}

.sidebar {
    display: none;
    width: 240px;
    background: white;
    border-right: 1px solid #e5e7eb;
    position: sticky; /* 固钉效果 */
    padding: 1rem;
    top: 10vh;
    height: 80vh;
}

.main-content {
    flex: 1;
    padding: 1.5rem;
    max-width: 1200px;
    margin: 0 auto;
}

.top-bar {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    margin-bottom: 1.5rem;
}

.search-input {
    max-width: 480px;
    flex: 1;
}

.view-mode-toggle {
    display: flex;
    gap: 0.5rem;
    align-items: center;
}


.mobile-categories {
    margin-bottom: 1rem;
    overflow: hidden;
    position: relative;
}

.categories-scroll {
    display: flex;
    overflow-x: auto;
    gap: 0.5rem;
    padding: 0.5rem 0;
    -webkit-overflow-scrolling: touch;
    scrollbar-width: none;
    /* Firefox */
    -ms-overflow-style: none;
    /* IE and Edge */
}

.categories-scroll::-webkit-scrollbar {
    display: none;
    /* Chrome, Safari and Opera */
}

.category-tag {
    flex-shrink: 0;
    cursor: pointer;
    padding: 6px 12px;
    font-size: 14px;
}

.content-list {
    display: flex;
    flex-direction: column;
    gap: 1rem;
}

.content-card {
    transition: all 0.3s ease;
}

.content-main {
    display: flex;
    flex-direction: column;
    gap: 1rem;
}

.content-title {
    font-size: 1.25rem;
    font-weight: 600;
    text-align: start;
    color: #1a1a1a;
    margin-bottom: 0.5rem;
    cursor: pointer;
}

.content-title:hover {
    color: var(--td-brand-color);
}

.content-description {
    color: #666;
    margin-bottom: 1rem;
    font-size: 0.875rem;
    text-align: start;
}

.content-meta {
    display: flex;
    align-items: center;
    gap: 1rem;
    color: #666;
    font-size: 0.875rem;
    flex-wrap: wrap;
}

@media (min-width: 768px) {
    .sidebar {
        display: block;
    }

    .mobile-categories {
        display: none;
    }

    .top-bar {
        flex-direction: row;
        align-items: center;
    }

}
@media (max-width: 768px) {
  .main-content {
    width: 85%;
    padding: 16px;
  }
  /* 分类标签优化 */
  .mobile-categories {
    margin: 0 -16px 1rem;
    padding: 0 16px;
  }

  .categories-scroll {
    padding-bottom: 8px; /* 增加滚动空间 */
  }

  .category-tag {
    font-size: 15px;
    padding: 8px 16px;
    border-radius: 20px;
    transition: all 0.2s;
  }

  /* 内容卡片优化 */
  .content-list {
    gap: 12px;
  }

  .content-card {
    margin: 0 -16px;
    width: calc(100% + 32px);
    border-radius: 0;
    box-shadow: none;
    border-bottom: 1px solid #eee;
  }

  .content-card:first-child {
    border-top: 1px solid #eee;
  }

  .content-title {
    font-size: 17px;
    line-height: 1.4;
  }

  .content-description {
    font-size: 15px;
    line-height: 1.6;
    color: #444;
  }

  /* 元数据优化 */
  .content-meta {
    gap: 8px 12px;
    font-size: 13px;
  }

  .content-meta > span {
    display: inline-flex;
    align-items: center;
    gap: 4px;
  }

  /* 标签样式优化 */
  .t-badge {
    font-size: 12px;
    padding: 4px 8px;
  }

  /* 全局触摸优化 */
  .t-icon {
    font-size: 16px !important;
  }

  button, [role="button"] {
    touch-action: manipulation; /* 优化移动端触摸响应 */
  }
}

/* 横屏适配 */
@media (max-width: 768px) and (orientation: landscape) {
  .view-mode-toggle .t-button {
    flex: 1 1 30%;
  }
}
</style>
