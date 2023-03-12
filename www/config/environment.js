/* jshint node: true */

module.exports = function(environment) {
  var ENV = {
    modulePrefix: 'open-ethereum-pool',
    environment: environment,
    rootURL: '/ubiq/',
    locationType: 'hash',
    EmberENV: {
      FEATURES: {
        // Here you can enable experimental features on an ember canary build
        // e.g. 'with-controller': true
      }
    },

    APP: {
      // API host and port
      ApiUrl: '/ubiq/',

      PoolName: 'UBIQ SOLO',
      CompanyName: 'Mine to buy a cool car :-)',

      // HTTP mining endpoint
      HttpHost: 'http://pool.lamba.top',
      HttpPort: 18880,

      // Stratum mining endpoint
      StratumHost: 'pool.lamba.top',
      StratumPort: 18180,

      // Fee and payout details
      PoolFee: '0.8%',
      PayoutThreshold: '1.0',
      PayoutInterval: '30m',

      // For network hashrate (change for your favourite fork)
      BlockTime: 21.7,
      BlockReward: 1.5,
      Unit: 'UBQ',
    }
  };

  if (environment === 'development') {
  }

  if (environment === 'test') {
  }

  if (environment === 'production') {
  }

  return ENV;
};
