<script setup>

  import Machine from '../components/Machine.vue'
  import {onMounted, ref, onBeforeMount, watch} from 'vue'
  import {NGrid, NGridItem, useNotification, useLoadingBar} from 'naive-ui'
  import {axios, baseAPIURL} from '../axios'

  import {storeToRefs} from 'pinia'
  import {useUserStore} from '../store'

  let notification = useNotification();
  const loadingBar = useLoadingBar()

  const {user: user} = storeToRefs(useUserStore())
  let machines = ref([])

  axios.get(baseAPIURL + "/users/new" , {
      headers: {
        "Authorization": ""
      }
    }).then((response) => {
        user.value = response.data
        loadingBar.finish()
        console.log(user.value)
        fetchData()
        setInterval(function(){
          fetchData()
        }, 3000)
    }).catch((error) => {
        loadingBar.finish()
      // loading.value = false
      // loadingBar.error()
      // notification.warning({title: "", content: error.code  + '\n' + error.message, duration: 1500});

    })

  let fetchData = () => {
    axios.get(baseAPIURL + "/machines/" , {
      headers: {
        "Authorization": user.value.Token
      }
    }).then((response) => {
        machines.value = response.data
        for (let index = 0; index < machines.value.length; index++) {
          machines.value[index].Key = machines.value[index].ID + new Date().getTime();
          
        }
        loadingBar.finish()
    }).catch((error) => {
        loadingBar.finish()

    })
  }


  

</script>

<template>
  <main >
    <h1>Vendor machine</h1>
    <br><br>
    <n-grid :cols="4" x-gap="12" item-responsive responsive="screen" style="width: 100%;">
      <n-grid-item v-for="machine in machines" v-bind:key="machine.Key"  >
        <Machine :machine="machine"/>
      </n-grid-item>
    </n-grid>
    
  </main>
</template>
