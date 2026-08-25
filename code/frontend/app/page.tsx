import { greetingResponse } from '../lib/mock/render-centered-hello-word';
import styles from './page.module.css';

export default function Home() {
  return (
    <main className={styles.screen} aria-label="Hello Word screen">
      <h1 className={styles.greeting}>{greetingResponse.data.greeting}</h1>
    </main>
  );
}
