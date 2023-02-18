import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
    vus: 3000,
    duration: '30s',
};

export default function () {
    const url = 'http://localhost:8080/message';
    const payload = JSON.stringify({
        text: 'I\'m k6!',
    });
    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    const res = http.post(url, payload, params);
    check(res, { 'status was 200': (r) => r.status == 200 });
    sleep(1);
}