<template>
  <div style="padding: 2rem">
    <ul class="container">
      <div class="row row-cols-4">
      <div class="col" style="margin-top: 3rem" v-for="book in books" :key="book.ID">
        <div class="card h-100 g-4" >
          <img style="height: 60%" v-bind:src='"data:image/png;base64," + book.Image' class="card-img-top">
          <div class="card-body">
            <h5 class="card-title">{{book.Name}}</h5>
            <p class="card-text">{{book.Description}}</p>
          </div>
          <div class="card-footer" >
            <button v-if="book.InCartCount === 0" v-on:click="addToCart(book.ID)" class="btn btn-primary">В корзину</button>
            <button v-if="book.InCartCount > 0" v-on:click="removeFromCart(book.ID)" class="btn btn-secondary">В корзине</button>
          </div>
        </div>
      </div>
      </div>
    </ul>
  </div>
</template>
<script>
export default {
  name: "Booklist",
  computed: {
    books() {
      return this.$store.state.books
    },
    session() {
      return this.$store.state.session
    },
  },
  methods: {
    addToCart(id) {
      this.$store.dispatch('addToCart', {session: this.session, id: id})
    },
    removeFromCart(id) {
      this.$store.dispatch('removeFromCart', {session: this.session, id: id})
    }
  }
}
</script>
<style scoped>
.book-list {
  padding: 0;
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
  align-items: stretch;
}
.book-list__item {
  list-style: none;
  width: 20%;
}
</style>