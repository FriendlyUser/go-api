<template>
  <div class="hello">
    <h1>Job Info</h1>
    <br>
    <br>
    <ul>
    </ul>
    <br>
    <hr>
    <br>
  </div>
</template>

<script>
import axios from 'axios';

const api = 'https://golang-job-api.herokuapp.com/api';

export default {
  name: 'hello',
  data: () => ({
    productName: null,
    productPrice: 0.0,
    products: [],
  }),

  methods: {
    /**
    async createProduct() {
      await axios.post(`${api}/jobs`, {
        name: this.productName,
        price: Number(this.productPrice),
      });

      // refresh the data
      this.retrieveProducts();
    },

    async deleteProduct(id) {
      // delete the product
      await axios.delete(`${api}/jobs/${id}`);

      // refresh the data
      const response = await axios.get(`${api}/products`);
      this.products = response.data;
    },
    */
    async retrieveProducts() {
      const response = await axios.get(`${api}/jobs`);
      this.products = response.data.sort((a, b) => a.id - b.id);
    },
  },

  async created() {
    this.retrieveProducts();
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
  padding-left: 40%;
}

li {
  /*display: inline-block;*/
  margin: 0 10px;
  text-align: left;
}

a {
  color: #42b983;
}
</style>
