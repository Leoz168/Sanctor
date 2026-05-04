import { Code2, GraduationCap, Home } from "lucide-react";

const highlights = [
  { icon: Code2, label: "Course help", detail: "Debugging, notes, and study groups" },
  { icon: Home, label: "Housing leads", detail: "Student-vetted rooms and sublets" },
  { icon: GraduationCap, label: "Campus life", detail: "Events, clubs, internships, and prep" },
];

interface CommunityAboutPanelProps {
  description: string;
}

export function CommunityAboutPanel({ description }: CommunityAboutPanelProps) {
  return (
    <section className="mt-8 rounded-[2rem] border border-gray-100 bg-white p-6 shadow-sm sm:p-8">
      <div className="grid gap-8 lg:grid-cols-[0.9fr_1.1fr]">
        <div>
          <p className="mb-3 text-xs font-black uppercase tracking-[0.2em] text-brand-orange/70">
            About the community
          </p>
          <h2 className="text-2xl font-bold text-gray-900">
            A practical hub for students who want signal, not noise.
          </h2>
          <p className="mt-4 text-base font-medium leading-8 text-gray-500">
            {description}
          </p>
        </div>

        <div className="grid gap-3">
          {highlights.map((item) => (
            <div
              key={item.label}
              className="flex items-center gap-4 rounded-2xl border border-orange-100 bg-orange-50/40 px-4 py-4"
            >
              <div className="flex h-11 w-11 shrink-0 items-center justify-center rounded-2xl bg-white text-brand-orange shadow-sm">
                <item.icon size={20} />
              </div>
              <div>
                <p className="font-bold text-gray-900">{item.label}</p>
                <p className="text-sm font-semibold text-gray-400">
                  {item.detail}
                </p>
              </div>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
}
