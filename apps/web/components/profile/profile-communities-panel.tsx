import Link from "next/link";
import { ProfileCommunityCard } from "@/components/profile/profile-community-card";
import { profilePanelClassName } from "@/components/profile/profile-panel-styles";

const activeCommunities = [
  {
    name: "Computer Science Collective",
    description:
      "A space for students to connect, share resources, and find their rhythm together.",
    image: "/images/community-1.jpg",
    href: "/communities/computer-science-collective",
    members: 1248,
    role: "Lead organizer",
    isLead: true,
  },
  {
    name: "Student Market & Swap",
    description:
      "Buy, sell, and trade the essentials that make student living easier.",
    image: "/images/community-5.jpg",
    href: "/communities/computer-science-collective",
    members: 3450,
    role: "Member",
  },
];

export function ProfileCommunitiesPanel() {
  return (
    <section className={`${profilePanelClassName} flex flex-col`}>
      <div className="flex flex-col gap-5 sm:flex-row sm:items-center sm:justify-between">
        <div>
          <p className="text-xs font-bold uppercase tracking-[0.22em] text-brand-orange">
            Active communities
          </p>
          <h1 className="mt-2 text-3xl font-bold tracking-tight text-gray-900">
            My Active Communities
          </h1>
          <p className="mt-2 text-sm font-semibold italic text-gray-500">
            Communities you are currently fostering.
          </p>
        </div>

        <Link
          href="/communities"
          className="inline-flex items-center justify-center rounded-2xl border border-brand-orange bg-white px-6 py-3 text-xs font-bold uppercase tracking-[0.16em] text-brand-orange shadow-sm shadow-orange-900/5 transition-all hover:bg-brand-cream active:scale-95"
        >
          Explore communities
        </Link>
      </div>

      <div className="mt-8 min-h-0 flex-1 overflow-y-auto px-3 pb-2">
        <div className="grid gap-8 lg:grid-cols-2">
          {activeCommunities.map((community) => (
            <ProfileCommunityCard key={community.name} {...community} />
          ))}
        </div>
      </div>
    </section>
  );
}
