<template>
    <div class="layout-container">
        <header class="header-fixed">
            <NavPC />
            <!-- 移动端导航先注释掉，需要时可以基于屏幕宽度动态显示 -->
            <!-- <NavmobileTop /> -->
            <!-- <Navmobile /> -->
        </header>

        <main class="main-content">
            <slot></slot>
        </main>
    </div>
</template>

<script setup>
import NavPC from '@/layout/components/navpc.vue';
import NavmobileTop from '@/layout/components/navmoblietop.vue';
import Navmobile from '@/layout/components/navmoblie.vue';
</script>

<style scoped>
.layout-container {
    /* min-height: 100vh; */
    display: flex;
    flex-direction: column;
}

.header-fixed {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    z-index: 100;
    background: #fff;
}

.main-content {
    margin-top: 64px; /* 与导航栏高度保持一致 */
    flex: 1;
    /* width: 100%; */
    padding: 20px;
}
</style>
