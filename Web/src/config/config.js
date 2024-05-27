import devConfig from './config.development.js';
import prodConfig from './config.production.js';

let config;

if (process.env.NODE_ENV === 'development') {
    config = devConfig;
} else if (process.env.NODE_ENV === 'production') {
    config = prodConfig;
} else {
    throw new Error('NODE_ENV not set or invalid');
}

export default config;
