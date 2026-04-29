import { Header } from "@/components/Header";

const features = [
  {
    title: "Communities",
    description: "Organize campus groups, membership, and moderation from one home base.",
  },
  {
    title: "Direct Messaging",
    description: "Keep peer-to-peer conversations close to the communities they support.",
  },
  {
    title: "Discovery",
    description: "Make it easy for students to find institutions, posts, and people that matter.",
  },
];

export default function HomePage() {
  return (
    <main className="shell">
      <Header />
      <section className="hero">
        <p className="eyebrow">Next.js Frontend</p>
        <h1>Sanctor is ready for a fresh frontend foundation.</h1>
        <p className="lede">
          This new app-router setup gives us a clean place to rebuild the product with server-rendered
          routes, typed components, and a simpler deployment path.
        </p>
      </section>
      <section className="feature-grid" aria-label="Platform features">
        {features.map((feature) => (
          <article className="feature-card" key={feature.title}>
            <h2>{feature.title}</h2>
            <p>{feature.description}</p>
          </article>
        ))}
      </section>
    </main>
  );
}
