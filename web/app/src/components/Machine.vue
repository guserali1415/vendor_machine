<script setup>
  let props = defineProps(["machine"])

  import {NCard, NButton, NIcon, useDialog, useLoadingBar, NSpace, useNotification} from 'naive-ui'
  import {onMounted, ref, onBeforeMount, watch} from 'vue'
  import {axios, baseAPIURL} from '../axios'

  import {ImageOutline} from '@vicons/ionicons5'
  import {DrinkCoffee16Filled, DrinkMargarita20Filled} from '@vicons/fluent'

  import {storeToRefs} from 'pinia'
  import {useUserStore} from '../store'


  const notification = useNotification()
  const dialog = useDialog()

  const {user: user} = storeToRefs(useUserStore())

  let loadingBar = useLoadingBar()

  let machine = ref(props.machine);



  function selectProduct(product){
      axios.get(baseAPIURL + "/machines/" + machine.value.ID + "/select_product/" + product.Name, {
        headers: {
          "Authorization": user.value.Token
        }
      }).then((response) => {
            dialog.warning({
              title: 'Confirm',
              content: 'Please insert coin to purchase this product?',
              positiveText: 'Insert coin',
              negativeText: 'Cancel',
              onPositiveClick: () => {
                purchaseProduct(product)
              },
              onNegativeClick: () => {
              }
            })
      }).catch((error) => {
          loadingBar.finish()
          notification.create({
            type: 'warning',
            content: 'Result',
            meta: '' + error.response.data.message,
            duration: 2500,
            keepAliveOnHover: true
          })
      })
  }

  function purchaseProduct(product){
    loadingBar.start()
      axios.get(baseAPIURL + "/machines/" + machine.value.ID + "/insert_coin/" + product.Price, {
        headers: {
          "Authorization": user.value.Token
        }
      }).then((response) => {
          loadingBar.finish()
          machine.value = response.data
          notification.create({
            type: 'success',
            content: 'Result',
            meta: product.Name + ' purchased successfully for ' + product.Price + " coins",
            duration: 2500,
            keepAliveOnHover: true
          })
      }).catch((error) => {
          loadingBar.finish()

      })
  }


  onMounted(() => {
  })

  

</script>

<template>
  <n-card :title="machine.Name" size="huge">
    <NSpace vertical>
      <n-button style="width: 100%" v-for="product in machine.Products" :key="product.ID" strong secondary type="primary" :disabled="product.Stock == 0"
        @click="selectProduct(product)">
        {{ product.Name}} 
        | Stock: 
        {{ product.Stock}}
        | Price: 
        {{ product.Price}}
        
        <template #icon>
          <n-icon v-if="product.Name == 'coffee'"><DrinkCoffee16Filled /></n-icon>
          <n-icon v-if="product.Name == 'drink'"><DrinkMargarita20Filled /></n-icon>

        </template>
      </n-button>
    </NSpace>
    

  </n-card>
</template>

<style scoped>
  .n-card {
    width: 100%;
  }
</style>
