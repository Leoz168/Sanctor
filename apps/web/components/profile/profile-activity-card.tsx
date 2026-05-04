import { Bookmark, Home, MessageSquare } from "lucide-react";

const activityItems = [
  {
    label: "Saved homes",
    value: "12",
    detail: "4 match your current filters",
    icon: Bookmark,
  },
  {
    label: "Community posts",
    value: "28",
    detail: "Computer Science Collective is most active",
    icon: MessageSquare,
  },
  {
    label: "Listings posted",
    value: "3",
    detail: "2 verified by community members",
    icon: Home,
  },
];

export function ProfileActivityCard() {
  return (
    <section className="grid gap-4 md:grid-cols-3">
      {activityItems.map((item) => {
        const Icon = item.icon;

        return (
          <article
            key={item.label}
            className="rounded-[1.5rem] border border-orange-100/70 bg-white/90 p-5 shadow-lg shadow-orange-900/5"
          >
            <div className="flex items-center justify-between gap-4">
              <div className="flex h-11 w-11 items-center justify-center rounded-2xl bg-orange-50 text-brand-orange">
                <Icon className="h-5 w-5" />
              </div>
              <span className="text-3xl font-bold text-gray-900">
                {item.value}
              </span>
            </div>
            <h3 className="mt-5 text-sm font-bold uppercase tracking-[0.18em] text-gray-400">
              {item.label}
            </h3>
            <p className="mt-2 text-sm font-semibold leading-relaxed text-gray-600">
              {item.detail}
            </p>
          </article>
        );
      })}
    </section>
  );
}
