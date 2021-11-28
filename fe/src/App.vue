<template>
  <div id="app">
    <Navbar />
    <router-view/>
  </div>
</template>

<script>
import Navbar from './components/Navbar'


export default {
  name: 'App',
  components: {
    Navbar,
  },
  computed: {
    session() {
      return this.$store.state.session
    },
    books() {
      return this.$store.state.books
    }
  },
  created() {
    console.log('created')
    console.log(this.$store)
    this.$store.dispatch('loadSession')
    this.$store.dispatch('loadBooks', this.session)
    this.$store.dispatch('loadCart', this.session)
  },

}
</script>

<style>
/*#app {*/
/*  font-family: Avenir, Helvetica, Arial, sans-serif;*/
/*  -webkit-font-smoothing: antialiased;*/
/*  -moz-osx-font-smoothing: grayscale;*/
/*  text-align: center;*/
/*  color: #2c3e50;*/
/*  margin-top: 60px;*/
/*}*/
</style>
