<template>
  <div>
    <div v-if="books.length === 0">Пустая корзина</div>
    <div v-if="books.length > 0">
      <div class="container" style="padding: 2rem">
        <div class="container" >
          <div class="row">
            <div class="col">
              <button class="btn-lg bg-primary" v-on:click="clearCart()" >Оформить заказ</button>
            </div>
            <div class="col">
              <div class="btn-lg bg-light" style="margin-left: 2rem;" disabled>
                {{'Общая сумма: ' + books.map((b) => b.Count * b.Price).reduce((a,b) => a + b)}}
              </div>
            </div>
          </div>
        </div>
        <ul class="container">
          <div class="row row-cols-5">
            <div class="col" style="margin-top: 3rem" v-for="book in books" :key="book.ID">
              <div class="card h-100 g-4" >
                <img style="height: 60%" v-bind:src='"data:image/png;base64," + book.Image' class="card-img-top">
                <div class="card-body">
                  <h5 class="card-title">{{book.Name}}</h5>
                  <p class="card-text">{{'Количество: ' + book.Count}}</p>
                </div>
              </div>
            </div>
          </div>
        </ul>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  name: 'Cart',
  computed: {
    books() {
      const cart = this.$store.state.cart
      if (cart) {
        return cart.Books
      }
      return []
    },
    session() {
      return this.$store.state.session
    },
  },
  methods: {
    clearCart() {
      this.$store.dispatch('clearCart', this.session)
      alert("Заказ успешно создан!")
    }
  }
}
</script>