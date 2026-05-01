import Image from "next/image";

export function CommunityHero() {
  return (
    <section className="px-4 py-12 sm:px-6 lg:px-8">
      <div className="mx-auto max-w-6xl">
        <div className="mb-10 text-center">
          <h1 className="mb-4 text-balance text-4xl font-bold text-foreground md:text-5xl">
            Computer Science Collective
          </h1>
          <p className="mx-auto max-w-2xl text-balance text-muted-foreground">
            A high-signal student space for course support, hackathon teams, housing leads, and everyday campus problem-solving.
          </p>
        </div>

        <div className="relative aspect-[16/9] overflow-hidden rounded-2xl md:aspect-[2/1]">
          <Image
            src="/images/community-1.jpg"
            alt="Computer Science Collective flagship community"
            fill
            className="object-cover"
            priority
          />
          <div className="absolute inset-0 bg-gradient-to-t from-black/60 via-transparent to-transparent" />

          <div className="absolute bottom-6 left-6">
            <span className="mb-3 inline-block rounded bg-primary px-3 py-1 text-xs font-semibold text-primary-foreground">
              FLAGSHIP COMMUNITY
            </span>
            <h2 className="text-2xl font-bold text-white md:text-3xl">
              U of T Student Housing + CS Network
            </h2>
          </div>
        </div>
      </div>
    </section>
  );
}
