DO $$
    BEGIN
        ALTER TABLE xreregistracion add column utm_source varchar DEFAULT '';
        ALTER TABLE xreregistracion add column utm_medium varchar DEFAULT '';
        ALTER TABLE xreregistracion add column utm_term varchar DEFAULT '';
        ALTER TABLE xreregistracion add column utm_content varchar DEFAULT '';
        ALTER TABLE xreregistracion add column utm_campaign varchar DEFAULT '';
    END;
$$;