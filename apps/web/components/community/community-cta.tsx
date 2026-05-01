import { Button } from "@/components/ui/button";

export function CommunityCTA() {
  return (
    <section className="px-4 py-12 sm:px-6 lg:px-8">
      <div className="mx-auto max-w-6xl">
        <div className="flex flex-col items-center justify-between gap-6 rounded-2xl bg-slate-800 px-8 py-12 md:flex-row md:py-16">
          <div>
            <h2 className="mb-2 text-2xl font-bold text-white md:text-3xl">
              Ready to join this community?
            </h2>
            <p className="max-w-md text-slate-400">
              Get access to student-run listings, discussion threads, and moderator-verified updates inside Sanctor.
            </p>
          </div>
          <Button className="whitespace-nowrap bg-primary px-8 py-6 text-base text-primary-foreground hover:bg-primary/90">
            Join Community
          </Button>
        </div>
      </div>
    </section>
  );
}
