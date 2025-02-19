<template>
    <div class="nav-top" v-if="!isMobile">
        <!-- 使用v-show或v-if根据isMobile的值显示或隐藏导航 -->
        <t-head-menu v-model="menu1Value" theme="light" @change="changeHandler" >
            <template #logo>
                <img height="44" src="@/assets/images/logo.jpg" alt="logo" />
                <span style="font-weight: 550;font-size: 32px;margin-left: 10px;">EduPal</span>
            </template>
            <t-menu-item value="item1">首页</t-menu-item>
            <t-menu-item value="item2">题库大厅</t-menu-item>
            <t-menu-item value="item3">AI刷题</t-menu-item>
            <template #operations>
                <t-button variant="text" shape="square">
                    <template #icon><t-icon name="search" /></template>
                </t-button>
                <t-button variant="text" shape="square">
                    <template #icon><t-icon name="notification" /></template>
                </t-button>

                <t-button variant="text" shape="square" @click="goToUserProfile">
                    <t-avatar
                        image="https://p26-passport.byteacctimg.com/img/user-avatar/424e10bd281f7bd137cd259593f07f26~140x140.awebp"
                        :hide-on-load-failed="false" />
                </t-button>
            </template>
        </t-head-menu>

    </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, watchEffect } from 'vue';
import { useRouter, useRoute } from 'vue-router';

const router = useRouter();
const route = useRoute();
const menu1Value = ref('item1');
const isMobile = ref(false);

// 定义函数来检查屏幕宽度
const checkScreenWidth = () => {
    isMobile.value = window.innerWidth < 768;
};

// 使用 watchEffect 的清理函数
const stopWatch = watchEffect(() => {
    switch (route.path) {
        case '/':
            menu1Value.value = 'item1';
            break;
        case '/quesbank':
            menu1Value.value = 'item2';
            break;
        case '/ai':
            menu1Value.value = 'item3';
            break;
        default:
            menu1Value.value = '';
    }
});

// 在组件卸载时停止监听
onUnmounted(() => {
    stopWatch(); // 停止路由监听
    window.removeEventListener('resize', checkScreenWidth);
});

// 组件挂载时添加监听
onMounted(() => {
    checkScreenWidth();
    window.addEventListener('resize', checkScreenWidth);
});

const changeHandler = (active) => {
    switch (active) {
        case 'item1':
            router.push('/');
            break;
        case 'item2':
            router.push('/quesbank');
            break;
        case 'item3':
            router.push('/practice');
            break;
    }
};

const goToUserProfile = () => {
    menu1Value.value = '';
    router.push('/user');
};
</script>

<style lang="less" scoped>
/* 为操作按钮添加左侧间距 */
.t-button {
    margin-right: 20px;
    /* 或者根据你的设计需求调整 */
}

/* 为了防止最后一个按钮也有右边距，可以使用 :not(:last-child) 选择器 */
.t-button:not(:last-child) {
    margin-right: 20px;
}
.nav-top {
    max-width: 100vw;
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    z-index: 1000;
    height: 64px;
    background: #0052D9; /* TDesign 默认主题蓝 */
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
}
:deep(.t-head-menu) {
    background: #0052D9 !important;
    color: #fff;
    height: 64px;
    line-height: 64px;

    .t-menu__item {
        color: #fff !important;

        &:hover {
            background: rgba(255, 255, 255, 0.1) !important;
        }

        &.t-is-active {
            background: rgba(255, 255, 255, 0.2) !important;
            color: #fff !important;
        }
    }
}
</style>
