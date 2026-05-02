import { HeroSearch } from "@/components/home/hero-search";

export function HomeHero() {
  return (
    <section className="relative pt-20 pb-32 overflow-hidden px-4">
      <div className="absolute top-0 left-1/2 -translate-x-1/2 w-full max-w-7xl h-full -z-10 bg-[radial-gradient(circle_at_50%_0%,rgba(255,237,213,0.5),transparent_65%)] pointer-events-none" />

      <div className="max-w-4xl mx-auto text-center" id="hero-content">
        <div>
          <h1 className="text-5xl md:text-7xl font-bold tracking-tight text-gray-900 mb-6 leading-[1.1]">
            Welcome to <span className="text-brand-orange">Rentling!!!</span>
          </h1>
          <p className="text-xl md:text-2xl text-gray-600 mb-12 max-w-2xl mx-auto font-medium">
            Find the right housing for you! Explore campus communities and trusted listings.
          </p>
        </div>

        <HeroSearch />
      </div>
    </section>
  );
}
