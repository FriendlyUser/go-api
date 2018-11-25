<template>
  <div class="hello">
    <h1>Job Info</h1>
    <br>
    <br>
    <ul id="example-1">
      <li v-for="item in jobs">
        {{ item.id }}
      </li>
    </ul>
    <br>
    <hr>
    <br>
  </div>
</template>

<script>
import axios from 'axios';

const api = 'https://golang-job-api.herokuapp.com/api';
// import { VueGoodTable } from 'vue-good-table';

export default {
  name: 'hello',
  data: () => ({
    productName: null,
    productPrice: 0.0,
    jobs: [],
  }),
  // add to component
  // components: {
  //   VueGoodTable,
  // },

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
      const response = await axios.get(`${api}/jobs`);
      this.products = response.data;
    },
    */
    async retrieveProducts() {
      const response = await axios.get(`${api}/jobs`);
      console.log(response)
      this.jobs = response.data.sort((a, b) => a.id - b.id);
    },
  },

  async created() {
    console.log('wait')
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
