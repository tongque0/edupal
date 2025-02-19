<template>
    <div class="top-nav" v-if="isMobile">
        <div class="nav-left">
            <img src="@/assets/images/logo.jpg" style="height: 44px;" alt="Logo" class="logo">
            <t-dropdown :options="menuOptions" @click="handleMenuClick">
                <t-button variant="text" style="color: #1e80ff; font-weight: 550;font-size: medium;">
                    {{ menuOptions.find(item => item.value === currentRoute)?.content || '首页' }}
                    <template #suffix>
                        <t-icon style="margin: -5px;" name="chevron-down" />
                    </template>
                </t-button>
            </t-dropdown>
        </div>

        <div class="nav-right">
            <t-button v-if="isRootOrSecondLevelRoute"
                variant="text"
                shape="circle"
                @click="goToUserProfile">
                <t-avatar
                    size="small"
                    image="https://p26-passport.byteacctimg.com/img/user-avatar/424e10bd281f7bd137cd259593f07f26~140x140.awebp"
                    :hide-on-load-failed="false" />
            </t-button>
            <t-button v-else
                variant="text"
                shape="circle"
                @click="goBack">
                <template #icon><t-icon name="arrow-left" /></template>
            </t-button>
        </div>
    </div>
</template>

<script setup>
import { onMounted, onUnmounted, ref, computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';

// 路由配置选项
const menuOptions = [
    {
        content: '首页',
        value: '/'
    },
    {
        content: '题库',
        value: '/quesbank'
    },
];


// 状态管理

const route = useRoute();
const router = useRouter();
// 路由层级判断
const isRootOrSecondLevelRoute = computed(() => {
    return route.matched.length <= 2;
});

// 判断当前路由
const currentRoute = computed(() => {
    return route.path;
});

// 屏幕宽度检测
const isMobile = ref(false);
const checkScreenWidth = () => {
    isMobile.value = window.innerWidth < 768;
};

// 生命周期钩子
onMounted(() => {
    checkScreenWidth();
    window.addEventListener('resize', checkScreenWidth);
});

onUnmounted(() => {
    window.removeEventListener('resize', checkScreenWidth);
});

// 事件处理方法
const handleMenuClick = (data) => {
    router.push(data.value);
};

const goToUserProfile = () => {
    router.push('/');
};

const goBack = () => {
    router.back();
};
</script>

<style scoped>
.top-nav {
    position: sticky;
    top: 0;
    z-index: 100;
    display: flex;
    justify-content: space-between;
    align-items: center;
    background: #ffffff;
    height: 56px;
    padding: 0 16px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
}

.nav-left {
    display: flex;
    align-items: center;
    /* gap: 8px; */
}

.nav-right {
    display: flex;
    align-items: center;
    gap: 8px;
    margin-right: 18px;
}

.logo {
    height: 32px;
    width: auto;
    object-fit: contain;
}

:deep(.t-dropdown) {
    font-size: 14px;
}

:deep(.t-button) {
    height: 32px;
    padding: 0 12px;
    font-size: 14px;
}

:deep(.t-dropdown__item) {
    padding: 10px 16px;
    font-size: 14px;
    color: #333;
}

:deep(.t-dropdown__item:hover) {
    background-color: #f5f5f5;
}

:deep(.t-avatar) {
    width: 32px;
    height: 32px;
}

:deep(.t-icon) {
    font-size: 16px;
}

/* 暗色模式适配 */
@media (prefers-color-scheme: dark) {
    .top-nav {
        background: #1a1a1a;
        box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2);
    }

    :deep(.t-dropdown__item) {
        color: #fff;
    }

    :deep(.t-dropdown__item:hover) {
        background-color: #2c2c2c;
    }
}

/* 移动端适配 */
@media screen and (max-width: 768px) {
    .top-nav {
        padding: 0 12px;
    }

    .logo {
        height: 28px;
    }

    :deep(.t-button) {
        padding: 0 8px;
    }
}
</style>
