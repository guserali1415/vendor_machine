<script setup>
import { RouterLink, RouterView, useRouter } from 'vue-router'
import { NLoadingBarProvider, NNotificationProvider, NDialogProvider, NConfigProvider, useLoadingBar } from 'naive-ui'
</script>

<template>
  <n-loading-bar-provider>
      <n-notification-provider>
        <NDialogProvider>
           <RouterView  />
        </NDialogProvider>
       
      </n-notification-provider>

    </n-loading-bar-provider>
</template>

<style scoped>

  h2, h3, p, small, img, span {
    -webkit-touch-callout: none; /* iOS Safari */
      -webkit-user-select: none; /* Safari */
      -khtml-user-select: none; /* Konqueror HTML */
        -moz-user-select: none; /* Old versions of Firefox */
          -ms-user-select: none; /* Internet Explorer/Edge */
              user-select: none; /* Non-prefixed version, currently
                                    supported by Chrome, Edge, Opera and Firefox */
  }



  :root{
    --blue-light: #AEDEF4;
    --blue-dark: #147eaf;
    --header-text-color: #202020;
    --grey-light: #f2f2f2;
  }
</style>
