import Link from "next/link";
import { ArrowLeft } from "lucide-react";
import { Field, inputClassName } from "@/components/create-post/field";
import { FormSection } from "@/components/create-post/form-section";
import { SegmentedControl } from "@/components/create-post/segmented-control";
import { Stepper } from "@/components/create-post/stepper";
import { ToggleRow } from "@/components/create-post/toggle-row";
import { SelectControl } from "@/components/forms/select-control";

export default function CreatePostPage() {
  return (
    <main className="min-h-screen bg-white px-4 py-8 font-sans text-[#1A1A1A]">
      <div className="mx-auto max-w-5xl">
        <Link
          href="/post-listings"
          className="mb-8 inline-flex items-center gap-2 text-sm font-bold text-gray-400 hover:text-brand-orange"
        >
          <ArrowLeft size={18} />
          Back to listings
        </Link>

        <header className="mb-12">
          <h1 className="text-4xl font-bold tracking-tight text-gray-900 sm:text-5xl">
            Create a <span className="text-brand-orange">Posting</span>
          </h1>
          <p className="mt-3 text-lg font-medium italic text-gray-500">
            Fill in the details below to find your perfect tenant or roommate.
          </p>
        </header>

        <form className="space-y-10">
          <FormSection title="Essential Information">
            <div className="space-y-6">
              <Field label="Listing Title">
                <input
                  className={inputClassName}
                  placeholder="e.g. Spacious Studio near St. George Campus"
                />
              </Field>

              <Field label="Address">
                <input className={inputClassName} placeholder="Full address of the property" />
              </Field>

              <div className="grid gap-6 md:grid-cols-2">
                <Field label="Price per Month ($)">
                  <input className={inputClassName} type="number" defaultValue={1000} />
                </Field>

                <Field label="Property Type">
                  <SelectControl
                    label="Property Type"
                    options={["Apartment", "House", "Studio", "Shared Room", "Dorm"]}
                    variant="create"
                  />
                </Field>
              </div>
            </div>
          </FormSection>

          <FormSection title="Room & Unit Specs">
            <div className="space-y-8">
              <div className="grid gap-6 md:grid-cols-3">
                <Stepper label="Total Rooms" initialValue={1} min={1} />
                <Stepper label="Rooms Occupied" initialValue={0} />
                <Stepper label="Bathrooms" initialValue={1} min={1} />
              </div>
              <ToggleRow />
            </div>
          </FormSection>

          <FormSection title="Social & Term Details">
            <div className="grid gap-8 lg:grid-cols-2">
              <SegmentedControl
                label="Gender Preference"
                options={["Coed", "Female Only", "Male Only"]}
                defaultValue="Coed"
              />
              <SegmentedControl
                label="Start Term"
                options={["Fall", "Spring", "Winter"]}
                defaultValue="Fall"
              />
            </div>
          </FormSection>

          <div className="space-y-7">
            <Field label="Quick Description (for card)">
              <textarea
                className={`${inputClassName} min-h-28 resize-y`}
                placeholder="Short summary of the housing..."
              />
            </Field>

            <Field label="Detailed Content">
              <textarea
                className={`${inputClassName} min-h-56 resize-y`}
                placeholder="Describe your property, rules, roommates, community vibes, etc..."
              />
            </Field>
          </div>

          <div className="mb-10 flex justify-end">
            <button className="rounded-2xl bg-brand-orange px-10 py-4 text-base font-black uppercase tracking-[0.18em] text-white shadow-xl shadow-brand-orange/25 transition-all hover:bg-orange-600 active:scale-[0.99]">
              Post My Listing
            </button>
          </div>
        </form>
      </div>
    </main>
  );
}
