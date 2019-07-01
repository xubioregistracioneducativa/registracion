DO $$
    BEGIN
        ALTER TABLE xreregistracion add column utm_source varchar;
        ALTER TABLE xreregistracion add column utm_medium varchar;
        ALTER TABLE xreregistracion add column utm_term varchar;
        ALTER TABLE xreregistracion add column utm_content varchar;
        ALTER TABLE xreregistracion add column utm_campaign varchar;
    END;
$$;